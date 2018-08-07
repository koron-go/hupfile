package hupfile

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

var (
	// ErrClosedAlready shows the File is closed already.
	ErrClosedAlready = errors.New("closed already")
)

// File is a write only file. It receive SIGHUP and re-open the file.
type File struct {
	name    string
	pidfile string
	sig     chan os.Signal
	closed  bool

	l sync.Mutex
	f *os.File
}

// New create a file to write.
func New(name, pidfile string) (*File, error) {
	f := &File{
		name:    name,
		pidfile: pidfile,
		sig:     make(chan os.Signal, 1),
	}
	// create PID file.
	if f.pidfile != "" {
		pf, err := os.Create(f.pidfile)
		if err != nil {
			return nil, err
		}
		_, err = pf.WriteString(strconv.Itoa(os.Getpid()))
		f.Close()
		if err != nil {
			return nil, err
		}
	}
	// start signal monitor.
	signal.Notify(f.sig, syscall.SIGHUP, os.Interrupt)
	go f.sigmonMain()
	return f, nil
}

// Write writes data to the underlying file.
func (f *File) Write(b []byte) (int, error) {
	f.l.Lock()
	defer f.l.Unlock()
	if f.closed {
		return 0, ErrClosedAlready
	}
	if f.f == nil {
		o, err := os.OpenFile(f.name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return 0, fmt.Errorf("failed to reopen file: %s", err)
		}
		f.f = o
	}
	return f.f.Write(b)
}

// Close close the file.
func (f *File) Close() error {
	f.l.Lock()
	defer f.l.Unlock()
	if f.closed {
		return nil
	}
	f.closed = true
	if f.f != nil {
		f.closeFile()
	}
	if f.sig != nil {
		signal.Stop(f.sig)
		close(f.sig)
		f.sig = nil
	}
	if f.pidfile != "" {
		os.Remove(f.pidfile)
	}
	return nil
}

// Reopen closes then openes the file with that name.
func (f *File) Reopen() error {
	f.l.Lock()
	defer f.l.Unlock()
	if f.closed {
		return ErrClosedAlready
	}
	if f.f == nil {
		return nil
	}
	f.closeFile()
	return nil
}

func (f *File) closeFile() {
	f.f.Sync()
	f.f.Close()
	f.f = nil
}

func (f *File) sigmonMain() {
	for {
		s, ok := <-f.sig
		if !ok {
			return
		}
		switch s {
		case syscall.SIGHUP:
			err := f.Reopen()
			if err != nil {
				return
			}
		case os.Interrupt:
			f.Close()
			return
		}
	}
}
