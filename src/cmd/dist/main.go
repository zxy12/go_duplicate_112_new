package main

import (
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	//zdebug.T("%v", "start build dist")
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

	gohostos = runtime.GOOS
	switch gohostos {
	// only set macos
	case "darwin":
		// Even on 64-bit platform, darwin uname -m prints i386.
		// We don't support any of the OS X versions that run on 32-bit-only hardware anymore.
		gohostarch = "amd64"
		// macOS 10.9 and later require clang
		defaultclang = true
	}
	sysinit()
	if gohostarch == "" {
		// Default Unix system.
		out := run("", CheckExit, "uname", "-m")
		switch {
		case strings.Contains(out, "x86_64"), strings.Contains(out, "amd64"):
			gohostarch = "amd64"
		case strings.Contains(out, "86"):
			gohostarch = "386"
		case strings.Contains(out, "arm"):
			gohostarch = "arm"
		case strings.Contains(out, "aarch64"):
			gohostarch = "arm64"
		case strings.Contains(out, "ppc64le"):
			gohostarch = "ppc64le"
		case strings.Contains(out, "ppc64"):
			gohostarch = "ppc64"
		case strings.Contains(out, "mips64"):
			gohostarch = "mips64"
			if elfIsLittleEndian(os.Args[0]) {
				gohostarch = "mips64le"
			}
		case strings.Contains(out, "mips"):
			gohostarch = "mips"
			if elfIsLittleEndian(os.Args[0]) {
				gohostarch = "mipsle"
			}
		case strings.Contains(out, "s390x"):
			gohostarch = "s390x"
		case gohostos == "darwin":
			if strings.Contains(run("", CheckExit, "uname", "-v"), "RELEASE_ARM_") {
				gohostarch = "arm"
			}
		default:
			fatalf("unknown architecture: %s", out)
		}
	}

	bginit()
	if len(os.Args) > 1 && os.Args[1] == "-check-goarm" {
		useVFPv1() // might fail with SIGILL
		println("VFPv1 OK.")
		useVFPv3() // might fail with SIGILL
		println("VFPv3 OK.")
		os.Exit(0)
	}
	xinit()
	xmain()
	xexit(0)
}
