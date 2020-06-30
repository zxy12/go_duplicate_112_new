package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	xprintf(`usage: go tool dist [command]
Commands are:

banner         print installation banner
bootstrap      rebuild everything
clean          deletes all built files
env [-p]       print environment (-p: include $PATH)
install [dir]  install individual directory
list [-json]   list all supported platforms
test [-h]      run Go test(s)
version        print Go version

All commands take -v flags to emit extra information.
`)
	xexit(2)
}

// The OS-specific main calls into the portable code here.
func xmain() {
	if len(os.Args) < 2 {
		usage()
	}
	cmd := os.Args[1]
	os.Args = os.Args[1:] // for flag parsing during cmd
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: go tool dist %s [options]\n", cmd)
		flag.PrintDefaults()
		os.Exit(2)
	}
	if f, ok := commands[cmd]; ok {
		f()
	} else {
		xprintf("unknown command %s\n", cmd)
		usage()
	}
}
