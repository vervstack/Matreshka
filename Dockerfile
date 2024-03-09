FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:latest as builder

WORKDIR /app
COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /deploy/server/matreshka-be ./cmd/matreshka-be/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine

WORKDIR /app
COPY --from=builder ./deploy/server/ .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 80

ENTRYPOINT ["./matreshka-be"]