ARG BASE_VERSION=10-slim
ARG BASE_DISTRO=debian
FROM ${BASE_DISTRO}:${BASE_VERSION}

LABEL description="Docker container for hugo static site generator to build CI - ${BASE_DISTRO} ${BASE_VERSION}"
LABEL maintainer="Karthik Kumar <support@malvahq.com>"

ENV TMP_DIRECTORY=/tmp
WORKDIR ${TMP_DIRECTORY}

RUN apt-get -y update && apt-get -y upgrade

ARG HUGO_VERSION=0.55.2
ENV HUGO_VERSION=${HUGO_VERSION}
ENV HUGO_FILE_PREFIX="hugo_extended_${HUGO_VERSION}"

ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz $TMP_DIRECTORY

RUN tar -xf $TMP_DIRECTORY/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz -C $TMP_DIRECTORY \
    && whoami \
    && mv $TMP_DIRECTORY/hugo /usr/local/bin \
    && ls -lt /usr/local/bin \
    && /usr/local/bin/hugo version \
    && rm $TMP_DIRECTORY/hugo*.tar.gz \
    && rm LICENSE README.md

WORKDIR /

EXPOSE 1313
