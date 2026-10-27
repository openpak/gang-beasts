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
| Port-base request | **22390** (auth 22390, secure 22391, healthz 22392; 22160 taken by super-mario-bros-35, 22200 requested by astroneer, first pick 22210 taken on prod by the Splatoon 3 service — moved at deploy 2026-10-27; verified live with `ss` on the host) |
| Game server ID | `2F56A000` — recovered 2026-10-27: managed const `NEX_SERVER_ID = 794206208` in `NintendoLobby`, consumed natively in `InitialisePiaSettings` (`0x2c4fc60/64/68` → `[x26+0x18]`) |
| Access key | `b9106845` — recovered 2026-10-27: managed const `NEX_ACCESS_KEY` in `NintendoLobby`, literal `0x6911A58`, same config struct `+0x20`; key-to-login link provisional pending CONNECT |
| LDN | `LOCAL_COMMUNICATION_ID = 72292340087201792`, wireless cryptokey `Boneloaf_GangBeasts_WirelessCryptoKey` (same class) |
| Registry status | **REGISTERED** 2026-10-27: `switch-nexd` entry `{"gang-beasts", 22390, "b9106845", ...}` + `bin/title-ports.json`; provenance in `switch-nex/docs/gang-beasts.md` |
| TODO-KEY | `CoatsinkServiceUrl` / `CoatsinkServiceFinalUrl` values — needs Il2CppDumper field read or asset dump |
| Open question | Which backend the Switch UI actually drives (NEX sessions vs Coatsink matchmaker vs Multiplay) — needs Il2CppDumper + Harbor runtime capture |

Note: an earlier gap line claimed NEX 4.6.7; no 4.6.7 string exists anywhere
in the binary. Serve 4.6.5.
