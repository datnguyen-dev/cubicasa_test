FROM alpine:latest

LABEL Author="Dat Nguyen <it.tandat@gmail.com>"

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

COPY dist/cubicasa_test /bin/
RUN chmod +x /bin/cubicasa_test
RUN mkdir -p /cubicasa_test
RUN mkdir -p /cubicasa_test/setup
WORKDIR /cubicasa_test

ADD ./setup/script.sql /cubicasa_test/setup/script.sql
VOLUME /cubicasa_test
EXPOSE 3100

CMD [ "cubicasa_test", "" ]
