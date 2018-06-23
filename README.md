# golem

[![GoDoc](https://godoc.org/github.com/ushios/golem?status.svg)](https://godoc.org/github.com/ushios/golem) [![CircleCI](https://circleci.com/gh/ushios/golem.svg?style=shield&circle-token=16878b14f171b0cd807f1ca57bde0fd6564ea1c5)](https://circleci.com/gh/ushios/golem) [![](https://dockerbuildbadges.quelltext.eu/status.svg?organization=ushios&repository=golem)](https://hub.docker.com/r/ushios/golem/builds/) [![Go Report Card](https://goreportcard.com/badge/github.com/ushios/golem)](https://goreportcard.com/report/github.com/ushios/golem) 

Get artifacts from CircleCI

## Usage

Download all artifacts to current directory.
```console
$ golem -vcs github -user ushios -repository golem -branch master
```

Get private repository's artifacts using token (using `-token` option).
```console
$ golem -token xxxxxxxx -vcs github -user ushios -repository golem -branch master
```

### options

```console
Usage of /usr/local/bin/golem:
  -branch string
    	branch name (default "master")
  -output string
    	output files destination directory (default "./")
  -prefix string
    	download file's prefix
  -repository string
    	respository name
  -token string
    	circle ci artifact build token
  -user string
    	github or bitbucket user name
  -vcs string
    	github or bitbucket (default "github")
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
