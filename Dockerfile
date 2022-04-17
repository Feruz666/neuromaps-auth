FROM golang:1.18.0-alpine3.14 as builder

WORKDIR /github.com/Feruz666/neuromaps-auth
COPY . .

RUN go install -mod=vendor

FROM alpine

RUN apk add --no-cache ca-certificates && \
  adduser -DH neuromaps-auth

COPY --from=builder /go/bin /usr/bin/
ADD config.yml .
ADD .env .
EXPOSE 8000 8000
ENV NMAUTH_CONFPATH=/etc/neuromaps-auth

ENTRYPOINT [ "/usr/bin/neuromaps-auth" ]