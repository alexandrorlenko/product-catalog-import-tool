FROM golang:1.14.2-alpine3.11 AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /src
RUN apk add --no-cache git
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates 
COPY . .
RUN go test ./...
RUN go get ./...
RUN go build -o /ts

# Final step
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /ts /src/service.yaml /
ENTRYPOINT ["/ts"]
