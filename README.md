# koron-go/hupfile

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/hupfile)](https://pkg.go.dev/github.com/koron-go/hupfile)
[![Actions/Go](https://github.com/koron-go/hupfile/workflows/Go/badge.svg)](https://github.com/koron-go/hupfile/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/hupfile)](https://goreportcard.com/report/github.com/koron-go/hupfile)

`hupfile` is a SIGHUP aware output only file which implement `io.WriteCloser`.
When did it accepts SIGHUP, it reopen an underlying file.
This is intended to use with logrotate.

## `hupredir` command

`hupredir` command reads from STDIN and write to a file.
And it accepts SIGHUP then reopen the file.
It makes your command support log rotation easily.

How to start `hupredir`:

```console
$ your-cmd-output-stdout | hupredir -out my.log -pid my.pid &
```

How to rotate log:

```console
$ mv my.log my.log.1
$ kill -HUP $(cat my.pid)
```

How to install and update `hupredir`:

```console
$ go install github.com/koron-go/hupfile/cmd/hupredir@latest
```

Or you can download the executable from [the release](https://github.com/koron-go/hupfile/releases/latest).
