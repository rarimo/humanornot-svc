FROM golang:1.19-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/humanornot-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/humanornot-svc /go/src/github.com/rarimo/humanornot-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/humanornot-svc /usr/local/bin/humanornot-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["humanornot-svc"]
