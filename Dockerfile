FROM golang:1.25-alpine AS builder

RUN apk add --no-cache bash gcc musl-dev linux-headers git

WORKDIR /go/src/github.com/guillembonet/bunetz
ADD . .
RUN go build -o ./build/bunetz-web ./cmd/main.go

FROM alpine:3.22

RUN apk update
RUN apk upgrade
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /go/src/github.com/guillembonet/bunetz/build/bunetz-web /usr/bin/bunetz-web

ENTRYPOINT ["/usr/bin/bunetz-web"]
