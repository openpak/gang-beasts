# Gang Beasts

OpenPak's Gang Beasts Server.

| | |
| --- | --- |
| Game | Gang Beasts |
| Network | Nintendo Switch |
| Title ID | `0100D58012E12000` |
| Licence | AGPL-3.0-only |

## What it serves

| Host | What | Served by |
| --- | --- | --- |
| NEX 4.6.5 (`*.s.n.srv.nintendo.net`) | Matchmaking sessions (MatchmakeExtension), NAT traversal, notifications | [`switch-nex`](https://github.com/openpak/switch-nex) after registration ([nex-spec.md](docs/nex-spec.md)) |
| Coatsink matchmaker / Multiplay | Matchmaker + dedicated-server legs | this repo (`gang-beastsd`, `/coatsink/*` stub) |
| `gbcostumedb-dev.elasticbeanstalk.com` | Costume database (plain HTTP) | this repo (`gang-beastsd`, `/costume/*` stub) |

## Run

```sh
go run ./cmd/gang-beastsd
```

| Variable | Meaning |
| --- | --- |
| `NX_INTERNAL_URL` | nx-baas internal API (default `http://127.0.0.1:20070`), used to check id tokens |
| `NX_INTERNAL_KEY` | its `X-Internal-Key` |

## Docs

- [Protocol](docs/protocol.md)
- [NEX registration request](docs/nex-spec.md)
- [Evidence](docs/evidence.md)
- [Ports](docs/ports.md)
- [Provenance](docs/provenance.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
