# golem

[![CircleCI](https://circleci.com/gh/ushios/golem.svg?style=svg&circle-token=16878b14f171b0cd807f1ca57bde0fd6564ea1c5)](https://circleci.com/gh/ushios/golem) [![](https://dockerbuildbadges.quelltext.eu/status.svg?organization=ushios&repository=golem)](https://hub.docker.com/r/ushios/golem/builds/)  

Get artifacts from CircleCI

## Usage

Download all artifacts to current directory
```console
$ golem -vcs github -user ushios -repository golem -branch master
```



## Development

Get dependencies.

```console
$ make dep
```

Testing

```console
$ make test
```
