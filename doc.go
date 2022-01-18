/*
Package hupfile provides SIGHUP aware output only file which implement
io.WriteCloser. When did it accepts SIGHUP, it reopen an underlying file. This
is intended to use with logrotate.
*/
package hupfile
