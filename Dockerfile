FROM golang:alpine as builder

WORKDIR /go/src/github.com/trennepohl/lru-cache-poc/

ENV GO111MODULE=on

RUN apk add git --no-cache

COPY . .

RUN go mod tidy && go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o lru ./cmd

FROM alpine:3.7

RUN addgroup -S lru && adduser -S -g lru lru

USER lru

COPY --from=builder /go/src/github.com/trennepohl/lru-cache-poc/lru /usr/bin/lru

EXPOSE 8080

CMD ["lru"]