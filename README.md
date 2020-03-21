# Docker Image for Hugo Static Site Generator

Dockerfile containing the `hugo` static site generator ( see https://gohugo.io )

The docker images are available at [codelios/hugo in Dockerhub](https://hub.docker.com/r/codelios/hugo).

If you are looking to work in hugo inside a docker image/container, go ahead , spin it as below and give it a try.


## Usage

### Alpine


```
docker pull codelios/hugo:0.65.0-alpine
```

### Debian distribution

```
docker pull codelios/hugo:0.65.0-debian
```

## Hugo Path

In the container image, Hugo is installed in `/usr/local/bin/hugo` .

Eg:

```
$ /usr/local/bin/hugo version
```
