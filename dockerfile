# Multi-stage layout
FROM golang:1.14 as builder

#ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Use distroless for prod, not for debug as it does not have ssh installed
COPY --from=builder /app/carrotautomation /app/
COPY --from=builder /app/migrations/ /app/migrations

WORKDIR /app
ENTRYPOINT ["./carrotautomation"]

