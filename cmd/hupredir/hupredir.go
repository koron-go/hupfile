package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/koron-go/hupfile"
)

var (
	flagOut string
	flagPid string
)

func main() {
	flag.StringVar(&flagOut, "out", "", "redirect destination file")
	flag.StringVar(&flagPid, "pid", "", "a file to record the process ID")
	flag.Parse()
	if flagOut == "" {
		log.Fatal("no destination file")
	}
	f, err := hupfile.New(flagOut, flagPid)
	if err != nil {
		log.Fatalf("failed to open destination file: %s", err)
	}
	io.Copy(f, os.Stdin)
	if err != nil && err != hupfile.ErrClosedAlready {
		log.Fatalf("failed to redirect: %s", err)
	}
}
