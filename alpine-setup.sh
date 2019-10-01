#!/bin/sh
apk update && apk upgrade
apk add --update libc6-compat libstdc++
