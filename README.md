# Docker Image for Hugo Static Site Generator

Dockerfile containing the `hugo` static site generator ( see https://gohugo.io )

The docker images are available at [codelios/hugo in Dockerhub](https://hub.docker.com/r/codelios/hugo).

If you are looking to work in hugo inside a docker image/container, go ahead , spin it as below and give it a try.


## Usage

### Alpine

```
docker pull codelios/hugo:latest-alpine
```

```
docker pull codelios/hugo:0.59.1-alpine
```

### Debian distribution

```
docker pull codelios/hugo:latest-debian
```

```
docker pull codelios/hugo:0.59.1-debian
```

## Hugo Path

Hugo is installed in `/usr/local/bin/hugo` .

Eg:

```
$ /usr/local/bin/hugo version
```
