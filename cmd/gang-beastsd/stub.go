package main

// Stubs for the two developer-side HTTP legs: the Coatsink matchmaker /
// Multiplay backend and the costume database. Neither contract is known yet —
// the recon recovered host/namespace facts but not paths or bodies (see
// docs/protocol.md) — so every call is logged with its body (capped) and
// answered 501 with a JSON error. A Harbor runtime capture against these
// stubs shows exactly what to implement next.

import (
	"io"
	"log"
	"net/http"
)

const maxLoggedBody = 64 << 10

// stub answers one call on the coatsink or costume legs.
func (s *server) stub(w http.ResponseWriter, r *http.Request, leg string) {
	n, _ := io.Copy(io.Discard, io.LimitReader(r.Body, maxLoggedBody))
	log.Printf("[gang-beasts] %s %s %s body=%d", leg, r.Method, r.URL.Path, n)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	io.WriteString(w, `{"error":"not-implemented"}`+"\n")
}
