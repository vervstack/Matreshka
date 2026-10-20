# ---- Go app build ----
FROM --platform=$BUILDPLATFORM golang:1.24.2-alpine AS builder

ARG VERSION=""

WORKDIR /src

RUN --mount=type=bind,target=/src,rw \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/go/mod \
    go mod download && \
    GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 \
    go build -ldflags "-X go.vervstack.ru/matreshka/internal/version.version=${VERSION}" \
        -o /deploy/server/service /src/cmd/service/main.go && \
    cp -r config /deploy/server/config && \
    mkdir -p /deploy/server/migrations && \
    cp -r /src/migrations/* /deploy/server/migrations/

FROM alpine:3.14

LABEL MATRESHKA_CONFIG_ENABLED=true

WORKDIR /app

COPY --from=builder /deploy/server/ .

RUN mkdir /app/data

EXPOSE 50049

VOLUME '/app/data'

ENTRYPOINT ["./service"]