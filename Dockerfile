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

FROM --platform=$BUILDPLATFORM golang AS builder

WORKDIR /app

COPY --from=webclient /dist /dist

RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 \
    go build -o /deploy/server/service ./cmd/service/main.go && \
    cp -r config /deploy/server/config && \
    if [ -d "./migrations" ];  then \
      cp -r ./migrations /deploy/server/migrations; \
    fi

FROM alpine

LABEL MATRESHKA_CONFIG_ENABLED=true

WORKDIR /app

COPY --from=builder /deploy/server/ .

RUN mkdir /app/data

EXPOSE 50049

VOLUME './data/matreshka-be.db'

ENTRYPOINT ["./service"]