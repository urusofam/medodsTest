FROM golang:1.24-alpine AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache git


COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o app ./main.go

FROM alpine:latest AS runner

RUN apk add --no-cache libc6-compat

COPY --from=builder /usr/local/src/app /

COPY .env .env

EXPOSE 8080

CMD ["/app"]