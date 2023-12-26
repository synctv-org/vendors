From alpine:latest as builder

ARG VERSION=dev

WORKDIR /vendors

COPY ./ ./

RUN apk add --no-cache bash curl gcc git go musl-dev 

RUN bash script/build.sh -v ${VERSION}

From alpine:latest

ENV PUID=0 PGID=0 UMASK=022

COPY --from=builder /vendors/build/vendors /usr/local/bin/vendors

COPY script/entrypoint.sh /entrypoint.sh

RUN apk add --no-cache bash ca-certificates su-exec tzdata && \
    rm -rf /var/cache/apk/* && \
    chmod +x /entrypoint.sh && \
    mkdir -p /vendors

WORKDIR /vendors

EXPOSE 9000

ENTRYPOINT [ "/entrypoint.sh" ]

CMD [ "server" ]