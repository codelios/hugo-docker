#!/bin/bash

function build_image() {
    IFS=' '
    read -ra argarray <<< $1
    BASE_DISTRO=${argarray[0]}
    BASE_VERSION=${argarray[1]}
    HUGO_VERSION=${argarray[2]}
    echo "inp: = $1, BASE_DISTRO = ${BASE_DISTRO}, BASE_VERSION=${BASE_VERSION}, HUGO_VERSION=${HUGO_VERSION}"
    docker build \
        --build-arg BASE_DISTRO=${BASE_DISTRO} \
        --build-arg BASE_VERSION=${BASE_VERSION} \
        --build-arg HUGO_VERSION=${HUGO_VERSION} \
        -t malvahq/ci-hugo:${HUGO_VERSION}-${BASE_DISTRO}${BASE_VERSION} .
}


RELEASES="
alpine 3.10 0.58.3
debian 10-slim 0.58.3
alpine 3.10 0.58.2
debian 10-slim 0.58.2
alpine 3.10 0.58.1
debian 10-slim 0.58.1
"

IFS=$'\n' ;for word in $RELEASES; do
    IFS=$'\n'
    build_image $word
done
