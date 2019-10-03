#!/bin/bash
if [ $# -ne 1 ]; then
    echo "Syntax: $0 <branch name>"
    exit 1
fi
sed -i 's/HUGO_VERSION=0.58.3/HUGO_VERSION=$1/' alpine/Dockerfile
sed -i 's/HUGO_VERSION=0.58.3/HUGO_VERSION=$1/' debian/Dockerfile
git com -m "feat: set hugo version $1" .
git push

