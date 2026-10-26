# Ports

| Port | Use |
| --- | --- |
| 23590 | plain HTTP: `/coatsink/*` and `/costume/*` stubs (behind the TLS terminator) |
| 23591 | `/healthz` |
| 23592–23599 | reserved |

NEX sessions use the `switch-nex` registry ports (requested base 22210, see
[nex-spec.md](nex-spec.md)); this service publishes none for them.
