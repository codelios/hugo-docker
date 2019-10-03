# Stage 1
# This stage merely downloads the binary. So the base distro and version can be hardcoded as we see fit
FROM alpine:3.10.1
RUN apk update && apk upgrade \
    && rm -vrf /var/cache/apk/*
ENV HUGO_VERSION=0.44
ENV HUGO_FILE_PREFIX="hugo_extended_${HUGO_VERSION}"
ENV TMP_DIRECTORY=/tmp
ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz $TMP_DIRECTORY
RUN tar -xf $TMP_DIRECTORY/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz -C $TMP_DIRECTORY \
    && rm $TMP_DIRECTORY/hugo*.tar.gz

# Final Stage
FROM alpine:3.10

LABEL description="Docker container for hugo static site generator - Alpine based"
LABEL vendor="MalvaHQ"

RUN apk update && apk upgrade \
    && apk add --update libc6-compat libstdc++ ca-certificates \
    && rm -vrf /var/cache/apk/*

COPY --from=0 /tmp/hugo /usr/local/bin

WORKDIR /

EXPOSE 1313
