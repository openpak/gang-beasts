# Provenance

Where each fact in [protocol.md](protocol.md) came from, per the
[clean-room policy](https://docs.openpak.org/contributors/clean-room/). Only the owner's own game
client was examined. No server code of the developer was seen or used.

| Fact | Source |
| --- | --- |
| Title ID `0100D58012E12000`, v0 base, Unity 2020.3.5f1 IL2CPP, SDK 11.4.x | NSP/NCA headers + `romfs/Data` layout, workspace `scratch/decompile-gang-beasts/RECON.md` |
| NEX 4.6 lib 4.6.5 (`S1140`, branch `NEX4.6`, pack `7923402a`) | version string literal, Switch main image (`scratch/decompile-gang-beasts/strings-main.txt`) |
| Transports (PRUDP/PRUDPLite/WebSocket/relay) + server URL + nncs templates | string literals, Switch main image |
| Protocol set present/absent | NEX RTTI in the Switch main image |
| NPLN-style login job names | string literals, Switch main image |
| UNET HLAPI lobby classes, `NET_*` message IDs, `NEX_SERVER_ID`/`NEX_ACCESS_KEY` field names (values unknown) | Unity managed metadata (`global-metadata.dat`, 8.5 MB) |
| Coatsink namespaces, `CoatsinkServiceUrl` fields, hardcoded UNET/Multiplay/cloud/costume-DB/coatsink URLs | Unity managed metadata |
| NPDM online services incl. `ldn:u` | `main.npdm.json`, workspace `scratch/decompile-gang-beasts/` |
| Absent middleware (PlayFab/EOS/Vivox/Demonware/GameSpy/Photon/UGS) | zero/false-positive-only hits across native + managed strings |

Full extraction notes: workspace `scratch/decompile-gang-beasts/RECON.md`
and `scratch/decompile-gang-beasts/findings.json`.
