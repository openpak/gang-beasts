# Gang Beasts — the client's contract

Read from the Switch executable (Gang Beasts `0100D58012E12000`, v0 base,
Unity 2020.3.5f1 IL2CPP), see [provenance](provenance.md). Keep this file in
step with `cmd/gang-beastsd`.

Every fact below comes from the workspace recon at
`scratch/decompile-gang-beasts/RECON.md` and
`scratch/decompile-gang-beasts/findings.json` (strings over the decompressed
NSO images plus the Unity managed metadata). No live traffic has been
captured yet; paths and bodies marked unknown await Il2CppDumper dumps and a
Harbor runtime capture.

## Engine

Unity 2020.3.5f1, IL2CPP (`global-metadata.dat` 8.5 MB, no managed DLLs).
NX SDK 11.4.3 libs, program NCA built with SDK 11.4.5.0. UNET HLAPI lobby
stack (`com.unity.multiplayer-hlapi.Runtime`): NetworkLobbyManager/Player,
NetworkDiscovery, NetworkMigrationManager host-migration, MatchMakingClient
with SecureTunnel. Lobby message IDs `NET_GANGS`, `NET_MEMBERS`,
`NET_PLAYERS`, `NET_PLATFORMS` (+`_CHAT`), `NETCHAT`, …

## NEX

NEX 4.6, library rev **4.6.5** (`NintendoSDK-NEX-for_NX-4_6_5-S1140-20201202-ja`,
branch `NEX4.6`, pack hash `7923402a`). Transports: PRUDP + PRUDPLite
(`prudp:/address=` builder, `prudps:/`), NEX WebSocket client transport, relay
(`Cannot connect to relay server`, `CutRouteOfRelayConnection`).

Protocols linked (RTTI): MatchMaking (+Ext), **MatchmakeExtension** sessions
(create/join/find/update/search, block list), NATTraversal (+Relay, Report),
Notification (+NintendoNotificationEvent), SecureConnection, TicketGranting,
AccountManagement, HealthService. NOT linked: Ranking, DataStore, Subscriber,
Utility, Shop, Messaging.

Login is NPLN-style: `JobBackEndServicesLogin`, `AcquireNsaIdToken` via
`SigloAccountFacade`, `JobSigloLogin`, `JobNNIDLogin`, `NgsBridgeImp`.
Server URL templates `g%08x-%%.s.n.srv.nintendo.net` and
`%s.%%.p.srv.nintendo.net`; NAT-check `nncs1/nncs2-%.n.n.srv.nintendo.net`.

NEX sessions are served by `switch-nex` once the title is registered — see
[nex-spec.md](nex-spec.md), which is this repo's registration request. The
`NEX_SERVER_ID` / `NEX_ACCESS_KEY` field names are present in the managed
metadata but their values are not recovered (needs Il2CppDumper).

## Coatsink backend (served here, stub)

The Switch porter (Coatsink) ships managed matchmaker/multiplay namespaces:
`Coatsink.Matchmaker.*` (Client/Server/Contracts/Data/Utils),
`Coatsink.Matchmaking.Utils`, `Coatsink.Multiplay.Client(.Qos/.Xsts)`,
`Coatsink.Multiplay.Server*`, with `CoatsinkServiceUrl` /
`CoatsinkServiceFinalUrl` fields. Recovered 2026-10-27 from the
`MatchmakingClientMonobehaviour` serialized fields (identical
180-byte object in `romfs/Data/level1` pathID 745 @`0x19c48` and
`romfs/Data/resources.assets` pathID 10536 @`0xe29188`):
`CoatsinkServiceUrl=https://matchmaking-service-prod.meattelecom.com`,
`FleetId=aa7c2580-5a72-11e7-9684-0242ac110009`,
`UnityProjectId=39848f19-bb83-4915-8516-cb4afbfcb371`,
`TitleName=Gangbeasts`. The exact FinalUrl template
(`MatchmakingClient.BuildCoatsinkServiceFinalUrl`, metadata method
51838 token `0x6000034`) and the live host choice (asset
`matchmaking-service-prod` vs literal `prod-matchmaker`) still need
a Harbor DNS/TLS capture.

Hardcoded endpoints in managed metadata:

- `https://mm.unet.unity3d.com` (UNET matchmaker)
- `https://qos.multiplay.com/v1/fleets/`
- `https://server-session.multiplay.com/v1/session`
- `https://cloud.connected.unity3d.com/`
- `https://coatsink.com/`

Whether the Switch build drives these at runtime (vs NEX sessions) is
unknown — the managed code paths need Il2CppDumper confirmation before
scoping shims. Until then `gang-beastsd` answers `/coatsink/*` with a logged
501 and JSON `{"error":"not-implemented"}`.

## Costume DB (served here, stub)

Costume database over **plain HTTP**: `http://gbcostumedb-dev.elasticbeanstalk.com/`
with new/update/get-bycode posts (exact paths not recovered). `gang-beastsd`
answers `/costume/*` with a logged 501 and JSON
`{"error":"not-implemented"}`.

## What is not here

- Matchmaking sessions: NEX 4.6.5 MatchmakeExtension, served by
  [`switch-nex`](https://github.com/openpak/switch-nex) after registration
  ([nex-spec.md](nex-spec.md)).
- NAT check: `nncs1/nncs2` hosts overlap OpenPak `nn-nncs`.
- Local wireless (`ldn:u` is in the NPDM service list) needs no server.
- Absent from the binary: PlayFab, EOS/Epic, Vivox, Demonware, GameSpy,
  Photon (hits are XR motion-to-photon), Unity Gaming Services
  (`Unity.Services` zero hits). "socket.io" hits are `nn::socket::Ioctl`;
  "dauth" hits are dormant Steam API names. USQP server-query strings are
  present but unscoped.
