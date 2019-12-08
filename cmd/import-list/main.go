package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/sanemat/go-importlist"
)

func main() {
	var data []byte
	var fi os.FileInfo
	var err error
	log.SetFlags(0)
	fi, err = os.Stdin.Stat()
	if err != nil {
		exitWithCode(err)
	}

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			exitWithCode(err)
		}
	}

	err = importlist.Run(os.Args[1:], data, os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp {
		exitWithCode(err)
	}
}

func exitWithCode(err error) {
	log.Println(err)
	exitCode := 1
	if ecoder, ok := err.(interface{ ExitCode() int }); ok {
		exitCode = ecoder.ExitCode()
	}
	os.Exit(exitCode)
}
