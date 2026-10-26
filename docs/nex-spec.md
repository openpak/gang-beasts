# NEX registration request — Gang Beasts (Switch)

Spec for the `switch-nex` per-title registry. This repo delivers the spec;
the registry edit itself is done serially later — do not register from here.

Source for every row: `scratch/decompile-gang-beasts/RECON.md` and
`scratch/decompile-gang-beasts/findings.json` (RTTI + strings over the
decompressed NSO images and Unity managed metadata). No live traffic yet.

| Item | Value |
| --- | --- |
| Title | Gang Beasts |
| Title ID | `0100D58012E12000` |
| Build | v0 base, Unity 2020.3.5f1 IL2CPP |
| NEX branch / lib | `NEX4.6` / **4.6.5** (`S1140 2020-12-02`, pack `7923402a`) |
| Transports | PRUDP, PRUDPLite, NEX WebSocket client, relay |
| Protocols to serve | MatchMaking (+Ext), MatchmakeExtension sessions (create/join/find/update/search, block list), NATTraversal (+Relay, Report), Notification (+NintendoNotificationEvent), SecureConnection, TicketGranting, AccountManagement, HealthService |
| Protocols NOT linked | Ranking, DataStore, Subscriber, Utility, Shop, Messaging |
| Auth shape | NPLN-style: `JobBackEndServicesLogin`, `AcquireNsaIdToken` via `SigloAccountFacade` |
| Server URL templates | `g%08x-%%.s.n.srv.nintendo.net`, `%s.%%.p.srv.nintendo.net` |
| Port-base request | **22210** (auth 22210, secure 22211, healthz 22212; 22160 was taken by super-mario-bros-35, 22200 requested by astroneer; verified against switch-nex `bin/title-ports.json` 2026-10-26) |
| TODO-KEY | `NEX_SERVER_ID` value — field name only in metadata, needs Il2CppDumper |
| TODO-KEY | `NEX_ACCESS_KEY` value — field name only in metadata, needs Il2CppDumper |
| TODO-KEY | `CoatsinkServiceUrl` / `CoatsinkServiceFinalUrl` values — needs Il2CppDumper or asset dump |
| Open question | Which backend the Switch UI actually drives (NEX sessions vs Coatsink matchmaker vs Multiplay) — needs Il2CppDumper + Harbor runtime capture |

Note: an earlier gap line claimed NEX 4.6.7; no 4.6.7 string exists anywhere
in the binary. Serve 4.6.5.
