FROM node:23-alpine3.20 AS webclient

WORKDIR /web

RUN --mount=type=bind,target=/web,rw \
    cd pkg/Matreshka-fe && \
    yarn && \
    yarn build &&\
    mv dist /dist

FROM --platform=$BUILDPLATFORM golang:1.23.4-alpine AS builder

WORKDIR /src

COPY --from=webclient /dist /dist

RUN --mount=type=bind,target=/src,rw \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    mkdir -p /src/internal/transport/web/dist && \
    mv /dist/* /src/internal/transport/web/dist && \
    GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 \
    go build -o /deploy/server/service /src/cmd/service/main.go && \
    cp -r config /deploy/server/config && \
    mkdir -p /deploy/server/migrations && \
    cp -r /src/migrations/* /deploy/server/migrations/

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /deploy/server/ .

EXPOSE 80

ENTRYPOINT ["./service"]
