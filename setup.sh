#!/bin/sh
if [ ! -f /etc/os-release ]; then
    echo "Unknown OS. Exit for now"
    exit 1
fi
grep alpine /etc/os-release
if [ $? -eq 0 ]; then
    $1/alpine-setup.sh
    exit 0
fi
grep Ubuntu /etc/os-release
if [ $? -eq 0 ]; then
    $1/debian-setup.sh
    exit 0
fi
grep Debian /etc/os-release
if [ $? -eq 0 ]; then
    $1/debian-setup.sh
    exit 0
fi
