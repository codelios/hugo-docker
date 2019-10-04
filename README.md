# Docker Image for Hugo Static Site Generator

This docker image contains the `hugo` static site generator ( see https://gohugo.io )

## Usage

### Alpine

```
docker pull malvahq/hugo:latest-alpine
```

```
docker pull malvahq/hugo:0.58.3-alpine
```

### Debian distribution

```
docker pull malvahq/hugo:latest-debian
```

```
docker pull malvahq/hugo:0.58.3-debian
```

## Hugo Path

Hugo is installed in `/usr/local/bin/hugo` .

Eg:

```
$ /usr/local/bin/hugo version
```
