FROM golang as builder
LABEL Config=matreshka
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deploy/server/matreshka-be ./cmd/matreshka-be/main.go

FROM alpine

WORKDIR /app
COPY --from=builder ./deploy/server/ .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 50051
EXPOSE 50052
ENTRYPOINT ["./matreshka-be"]