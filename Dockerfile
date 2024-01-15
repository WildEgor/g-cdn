# Base Stage
FROM golang:1.21-alpine as base
COPY . /app/
WORKDIR /app/
RUN go mod download && mkdir -p dist

# Development Stage
FROM golang:1.21-alpine as dev
WORKDIR /app
RUN go install -mod=mod github.com/cosmtrek/air
CMD ["air", "-c", ".air-unix.toml"]

# # Test Stage
# FROM base as test
# ENTRYPOINT make test

# Build Production Stage
FROM base as builder
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/app cmd/main.go

# Production Stage
FROM alpine:latest as production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/dist/app .
EXPOSE 8888
CMD ["./app"]