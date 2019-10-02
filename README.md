# Docker Image for Hugo Static Site Generator

This docker image contains the `hugo` executable, representing the static site generator available at https://gohugo.io .

## Usage

### Alpine
```
docker pull malvahq/ci-hugo:0.58.3-alpine3.10
```

### Debian distribution
```
docker pull malvahq/ci-hugo:0.58.3-debian10-slim
```

## Hugo Path

Hugo is installed in `/usr/local/bin/hugo` .

Eg:

```
$ /usr/local/bin/hugo version
```
