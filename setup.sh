#!/bin/sh
if [ $# -ne 1 ]; then
    echo "Syntax: $0 <current dir>"
    exit 1
fi
if [ ! -f /etc/os-release ]; then
    # TODO: Possibly read /etc/issue in case of CentOS. But for now, not supported yet.
    echo "Unknown OS. Exit for now"
    exit 2
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
