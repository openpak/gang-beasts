package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeIDs map[string]error

func (f fakeIDs) check(tok string) error {
	if err, ok := f[tok]; ok {
		return err
	}
	return errTokenRejected
}

func call(t *testing.T, s *server, method, target, body string) *httptest.ResponseRecorder {
	t.Helper()
	var rdr *strings.Reader
	if body == "" {
		rdr = strings.NewReader("")
	} else {
		rdr = strings.NewReader(body)
	}
	w := httptest.NewRecorder()
	s.ServeHTTP(w, httptest.NewRequest(method, target, rdr))
	return w
}

func TestHealthz(t *testing.T) {
	w := call(t, &server{ids: fakeIDs{}}, http.MethodGet, "/healthz", "")
	if w.Code != 200 || w.Body.String() != "ok\n" {
		t.Errorf("healthz: got %d %q, want 200 %q", w.Code, w.Body.String(), "ok\n")
	}
}

func TestStubsAnswer501JSON(t *testing.T) {
	s := &server{ids: fakeIDs{}}
	for _, target := range []string{
		"/coatsink/session",
		"/coatsink/",
		"/costume/get-bycode",
		"/costume/",
	} {
		w := call(t, s, http.MethodPost, target, `{"x":1}`)
		if w.Code != http.StatusNotImplemented {
			t.Errorf("%s: got %d, want 501", target, w.Code)
		}
		if ct := w.Header().Get("Content-Type"); ct != "application/json" {
			t.Errorf("%s: content-type %q, want application/json", target, ct)
		}
		if !strings.Contains(w.Body.String(), "not-implemented") {
			t.Errorf("%s: body %q missing not-implemented", target, w.Body.String())
		}
	}
}

func TestUnknownIs404(t *testing.T) {
	w := call(t, &server{ids: fakeIDs{}}, http.MethodGet, "/nope", "")
	if w.Code != http.StatusNotFound {
		t.Fatalf("unknown path: %d", w.Code)
	}
}
