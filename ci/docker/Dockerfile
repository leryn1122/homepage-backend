ARG GO_VERSION=1.19
ARG DEBIAN_VERSION=11.6-slim
FROM golang:${GO_VERSION}-alpine AS build

ARG MIRRORS_SOURCE="dl-cdn.alpinelinux.org" \
    GOPROXY="https://goproxy.io,direct"

WORKDIR /opt

RUN sed -i "s/dl-cdn.alpinelinux.org/${MIRRORS_SOURCE}/g" /etc/apk/repositories \
 && apk add --update --no-cache \
      make \
      binutils

ADD . /opt

RUN GO111MODULE=on \
    GOPROXY=${GOPROXY} \
    make install

RUN make build

FROM debian:${DEBIAN_VERSION} AS runtime

LABEL maintainers="Leryn <leryn1122@gmail.com>"
LABEL description="homepage"

COPY etc/config.toml etc/config.toml
COPY --from=build /opt/target/exe /exe

ENTRYPOINT [ "/exe" ]