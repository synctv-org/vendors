From alpine:latest as builder

WORKDIR /vendors

COPY ./ ./

RUN apk add --no-cache bash curl gcc git go musl-dev 

RUN make build

From alpine:latest

ENV PUID=0 PGID=0 UMASK=022

COPY --from=builder /vendors/bin/vendors /usr/local/bin/vendors

COPY script/entrypoint.sh /entrypoint.sh

RUN apk add --no-cache bash ca-certificates su-exec tzdata && \
    rm -rf /var/cache/apk/* && \
    chmod +x /entrypoint.sh && \
    mkdir -p /vendors

WORKDIR /vendors

EXPOSE 9000

ENTRYPOINT [ "/entrypoint.sh" ]

CMD [ "server" ]