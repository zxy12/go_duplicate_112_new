package main

import (
	"log"
	"os"

	"github.com/zxy12/go_duplicate_112_new/src/zdebug"
)

func main() {
	zdebug.T("%v", "start build dist")
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ltime)
	main1()
}

func main1() {
	os.Setenv("TERM", "dumb") // disable escape codes in clang errors

	// provide -check-armv6k first, before checking for $GOROOT so that
	// it is possible to run this check without having $GOROOT available.
	if len(os.Args) > 1 && os.Args[1] == "-check-armv6k" {
		useARMv6K() // might fail with SIGILL
		println("ARMv6K supported.")
		os.Exit(0)
	}
}
