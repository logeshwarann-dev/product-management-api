ARG GO_VERSION=1.24.0

# Build stage
FROM golang:${GO_VERSION}-alpine AS builder

ARG APP 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

RUN go build -ldflags="-s -w" -o app ./cmd/${APP}


# Runtime
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./app"]
