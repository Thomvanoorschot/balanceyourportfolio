##################
# Build Stage 1/3   -> dependencies
##################
FROM golang:1.21 AS dependencies

WORKDIR /src

COPY go.* ./

RUN --mount=type=cache,target=/go/pkg/mod/ go mod download -x

##################
# Build Stage 2/3   -> Build
##################
FROM dependencies AS Build

COPY . ./

RUN --mount=type=cache,target=/go/pkg/mod/ CGO_ENABLED=0 go build \
    -C cmd/api \
    -installsuffix 'static' \
    -ldflags '-s -w' \
    -o /app


##################
# Build Stage 3/3   -> Final image
##################
FROM alpine:3.19 AS final

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.13 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe \

RUN adduser -D -g '' nonroot
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /app /app

ENTRYPOINT ["/app"]
