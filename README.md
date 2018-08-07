# koron-go/hupfile

[![GoDoc](https://godoc.org/github.com/koron-go/hupfile?status.svg)](https://godoc.org/github.com/koron-go/hupfile)
[![CircleCI](https://img.shields.io/circleci/project/github/koron-go/hupfile/master.svg)](https://circleci.com/gh/koron-go/hupfile/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/hupfile)](https://goreportcard.com/report/github.com/koron-go/hupfile)

`hupfile` is a SIGHUP aware output only file which implement `io.WriteCloser`.
When did it accepts SIGHUP, it reopen an underlying file.
This is intended to use with logrotate.

## hupredir

`hupredir` read from STDIN and write to a file.
And it accepts SIGHUP then reopen the file.
It makes your command support log rotation easily.

How to start `hupredir`:

```console
$ your-cmd-output-stdout | hupredir -out my.log -pid my.pid &
```

How to rotate log:

```console
$ mv my.log my.log.1
$ kill -HUP `cat my.pid`
```

How to install and update `hupredir`:

```console
$ go get -u github.com/koron-go/hupfile/cmd/hupredir
```
