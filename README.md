# Docker Image for Hugo Static Site Generator

Dockerfile containing the `hugo` static site generator ( see https://gohugo.io )

The docker images are available at [malvahq/hugo in Dockerhub](https://hub.docker.com/r/malvahq/hugo).

If you are looking to work in hugo inside a docker image/container, go ahead , spin it as below and give it a try.


## Usage

### Alpine

```
docker pull malvahq/hugo:latest-alpine
```

```
docker pull malvahq/hugo:0.62.2-alpine
```

### Debian distribution

```
docker pull malvahq/hugo:latest-debian
```

```
docker pull malvahq/hugo:0.62.2-debian
```

## Hugo Path

Hugo is installed in `/usr/local/bin/hugo` .

Eg:

```
$ /usr/local/bin/hugo version
```
