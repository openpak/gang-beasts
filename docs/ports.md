# Ports

| Port | Use |
| --- | --- |
| 23560 | plain HTTP: `/coatsink/*` and `/costume/*` stubs (behind the TLS terminator) |
| 23561 | `/healthz` |
| 23562–23569 | reserved |

NEX sessions use the `switch-nex` registry ports (requested base 22160, see
[nex-spec.md](nex-spec.md)); this service publishes none for them.
