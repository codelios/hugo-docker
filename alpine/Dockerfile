ARG BASE_VERSION=3.10
ARG BASE_DISTRO=alpine
FROM ${BASE_DISTRO}:${BASE_VERSION}

LABEL description="Docker container for hugo static site generator - Alpine based"
LABEL maintainer="Karthik Kumar <support@malvahq.com>"

ENV TMP_DIRECTORY=/tmp
WORKDIR ${TMP_DIRECTORY}

RUN apk update && apk upgrade \
    && apk add --update libc6-compat libstdc++

ENV HUGO_VERSION=0.58.3
ENV HUGO_FILE_PREFIX="hugo_extended_${HUGO_VERSION}"


ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz $TMP_DIRECTORY

RUN tar -xf $TMP_DIRECTORY/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz -C $TMP_DIRECTORY \
    && mv $TMP_DIRECTORY/hugo /usr/local/bin \
    && /usr/local/bin/hugo version \
    && rm $TMP_DIRECTORY/hugo*.tar.gz \
    && rm LICENSE README.md

WORKDIR /

EXPOSE 1313
