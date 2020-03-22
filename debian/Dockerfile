# Stage 1
# This stage merely downloads the binary. So the base distro and version can be hardcoded as we see fit
FROM debian:10-slim
ENV HUGO_VERSION=0.55.1
ENV HUGO_FILE_PREFIX="hugo_${HUGO_VERSION}"
ENV TMP_DIRECTORY=/tmp
ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz $TMP_DIRECTORY
RUN tar -xf $TMP_DIRECTORY/${HUGO_FILE_PREFIX}_Linux-64bit.tar.gz -C $TMP_DIRECTORY \
    && rm $TMP_DIRECTORY/hugo*.tar.gz

# Final Stage
FROM debian:10-slim

LABEL description="Docker image for hugo static site generator - Debian based"
LABEL vendor="MalvaHQ"

RUN apt-get -y update \
    && apt-get -y upgrade \
    && apt-get -y install ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=0 /tmp/hugo /usr/local/bin

WORKDIR /

EXPOSE 1313
