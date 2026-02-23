FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o messenger ./cmd/messenger

FROM alpine:3.19
COPY --from=builder /app/messenger /messenger
CMD ["/messenger"]

EXPOSE 8080