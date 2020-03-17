# Docker Image for Hugo Static Site Generator

Dockerfile containing the `hugo` static site generator ( see https://gohugo.io )

The docker images are available at [feldci/hugo in Dockerhub](https://hub.docker.com/r/feldci/hugo).

If you are looking to work in hugo inside a docker image/container, go ahead , spin it as below and give it a try.


## Usage

### Alpine

```
docker pull feldci/hugo:latest-alpine
```

```
docker pull feldci/hugo:0.67.0-alpine
```

### Debian distribution

```
docker pull feldci/hugo:latest-debian
```

```
docker pull feldci/hugo:0.67.0-debian
```

## Hugo Path

In the container image, Hugo is installed in `/usr/local/bin/hugo` .

Eg:

```
$ /usr/local/bin/hugo version
```
