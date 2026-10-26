// Command gang-beastsd answers Gang Beasts' developer-side HTTP surface for the
// Switch build (0100D58012E12000, Unity 2020.3.5f1 IL2CPP) when a console is
// redirected to OpenPak.
//
// Matchmaking sessions themselves are NEX 4.6.5 (MatchmakeExtension sessions,
// served by platforms/switch/games/switch-nex once the title is registered —
// see docs/nex-spec.md), and the lobby is UNET HLAPI. This service is the
// rest of the stack the client reaches over HTTP:
//
//	Coatsink matchmaker / Multiplay legs (Coatsink.Matchmaker.*,
//	  Coatsink.Multiplay.*)           /coatsink/* stub: logged, 501 JSON
//	Costume DB (gbcostumedb-dev...)    /costume/* stub: logged, 501 JSON
//	Health                             /healthz: "ok"
//
// Both stubs log what the client asks and answer 501 with a JSON error, so a
// Harbor capture shows the exact paths and bodies to implement next. Routes
// it does not know are logged with their body and answered 404.
//
// It listens on plain HTTP behind the TLS terminator (the costume DB host is
// plain HTTP in the client already). Identity checks go to nx-baas, the same
// way unrailedd asks it.
//
// The contract was read out of the Switch executable; see docs/protocol.md.
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

var (
	addr       = flag.String("addr", ":23590", "plain HTTP listen address (behind the TLS terminator)")
	healthAddr = flag.String("health-addr", ":23591", "health listen address")
)

func main() {
	flag.Parse()
	s := &server{ids: newIdentityClient()}

	go func() {
		h := http.NewServeMux()
		h.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { io.WriteString(w, "ok\n") })
		log.Printf("[gang-beasts] health on %s", *healthAddr)
		log.Fatal(http.ListenAndServe(*healthAddr, h))
	}()
	log.Printf("[gang-beasts] serving HTTP on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, s))
}

type server struct {
	ids identityChecker
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/healthz":
		io.WriteString(w, "ok\n")
	case r.URL.Path == "/coatsink" || hasPrefix(r.URL.Path, "/coatsink/"):
		s.stub(w, r, "coatsink")
	case r.URL.Path == "/costume" || hasPrefix(r.URL.Path, "/costume/"):
		s.stub(w, r, "costume")
	default:
		log.Printf("[gang-beasts] !! unhandled %s %s", r.Method, r.URL.Path)
		http.NotFound(w, r)
	}
}

func hasPrefix(s, p string) bool {
	return len(s) >= len(p) && s[:len(p)] == p
}
