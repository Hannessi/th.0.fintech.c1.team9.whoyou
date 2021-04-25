from alpine:latest

RUN mkdir -p /apps/whoYou

COPY assembly /apps/whoYou
COPY config.json /apps/whoYou

ENTRYPOINT /apps/whoYou/assembly -config-file /apps/whoYou/config.json
