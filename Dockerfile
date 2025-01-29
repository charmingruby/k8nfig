# STEP 1: Build
FROM golang:1.23.3-alpine AS builder
RUN apk update && apk add --no-cache git upx
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH="amd64" go build -ldflags="-w -s" -o ./bin/k8nfig ./cmd/k8nfig/main.go

# STEP 2: Optimize
RUN upx --best --lzma ./bin/k8nfig

# STEP 3: Run
FROM alpine:latest
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
COPY --from=builder /build/bin/k8nfig ./k8nfig
EXPOSE 8080
CMD [ "/app/k8nfig" ]