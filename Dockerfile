# syntax=docker/dockerfile:1
# Gang Beasts' developer-side HTTP surface (Coatsink matchmaker/Multiplay and
# costume DB stubs) for the Switch build. Plain HTTP on 23560 behind the TLS
# terminator; health on 23561.
FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build
ARG TARGETOS TARGETARCH
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GOWORK=off go build -trimpath -buildvcs=false -ldflags="-s -w" -o /out/gang-beastsd ./cmd/gang-beastsd

FROM gcr.io/distroless/static
COPY --from=build /out/gang-beastsd /gang-beastsd
USER nonroot:nonroot
EXPOSE 23560 23561
ENTRYPOINT ["/gang-beastsd"]
