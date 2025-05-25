FROM node:23-alpine3.20 AS webclient

WORKDIR /web


RUN --mount=type=bind,target=/web,rw \
# Step 1: Build the API lib
    cd /web/pkg/web/@vervstack/matreshka && \
    yarn && \
    yarn build && \
# Step 2: Install and build Vue app (now that web is built)
    cd /web/pkg/web/Matreshka-UI && \
    yarn && \
    yarn build && \
    mv dist /dist

FROM --platform=$BUILDPLATFORM golang:1.24.2-alpine AS builder

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

RUN mkdir /app/data

EXPOSE 50049
ENTRYPOINT ["./service"]
