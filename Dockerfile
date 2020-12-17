FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

COPY bin/auth-microservice /usr/local/bin/auth-microservice

CMD ["/usr/local/bin/auth-microservice"]
