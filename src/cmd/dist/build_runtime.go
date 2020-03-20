// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

/*
 * Helpers for building runtime.
 */

// mkzversion writes zversion.go:
//
//  package sys
//
//  const TheVersion = <version>
//  const Goexperiment = <goexperiment>
//  const StackGuardMultiplier = <multiplier value>
//
func mkzversion(dir, file string) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated by go tool dist; DO NOT EDIT.\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "package sys\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "const TheVersion = `%s`\n", findgoversion())
	fmt.Fprintf(&buf, "const Goexperiment = `%s`\n", os.Getenv("GOEXPERIMENT"))
	fmt.Fprintf(&buf, "const StackGuardMultiplierDefault = %d\n", stackGuardMultiplierDefault())

	writefile(buf.String(), file, writeSkipSame)
}

// mkzbootstrap writes cmd/internal/objabi/zbootstrap.go:
//
//  package objabi
//
//  const defaultGOROOT = <goroot>
//  const defaultGO386 = <go386>
//  const defaultGOARM = <goarm>
//  const defaultGOMIPS = <gomips>
//  const defaultGOMIPS64 = <gomips64>
//  const defaultGOOS = runtime.GOOS
//  const defaultGOARCH = runtime.GOARCH
//  const defaultGO_EXTLINK_ENABLED = <goextlinkenabled>
//  const version = <version>
//  const stackGuardMultiplierDefault = <multiplier value>
//  const goexperiment = <goexperiment>
//
// The use of runtime.GOOS and runtime.GOARCH makes sure that
// a cross-compiled compiler expects to compile for its own target
// system. That is, if on a Mac you do:
//
//  GOOS=linux GOARCH=ppc64 go build cmd/compile
//
// the resulting compiler will default to generating linux/ppc64 object files.
// This is more useful than having it default to generating objects for the
// original target (in this example, a Mac).
func mkzbootstrap(file string) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated by go tool dist; DO NOT EDIT.\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "package objabi\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "import \"runtime\"\n")
	fmt.Fprintln(&buf)
	fmt.Fprintf(&buf, "const defaultGO386 = `%s`\n", go386)
	fmt.Fprintf(&buf, "const defaultGOARM = `%s`\n", goarm)
	fmt.Fprintf(&buf, "const defaultGOMIPS = `%s`\n", gomips)
	fmt.Fprintf(&buf, "const defaultGOMIPS64 = `%s`\n", gomips64)
	fmt.Fprintf(&buf, "const defaultGOOS = runtime.GOOS\n")
	fmt.Fprintf(&buf, "const defaultGOARCH = runtime.GOARCH\n")
	fmt.Fprintf(&buf, "const defaultGO_EXTLINK_ENABLED = `%s`\n", goextlinkenabled)
	fmt.Fprintf(&buf, "const version = `%s`\n", findgoversion())
	fmt.Fprintf(&buf, "const stackGuardMultiplierDefault = %d\n", stackGuardMultiplierDefault())
	fmt.Fprintf(&buf, "const goexperiment = `%s`\n", os.Getenv("GOEXPERIMENT"))

	writefile(buf.String(), file, writeSkipSame)
}

// stackGuardMultiplierDefault returns a multiplier to apply to the default
// stack guard size. Larger multipliers are used for non-optimized
// builds that have larger stack frames.
func stackGuardMultiplierDefault() int {
	for _, s := range strings.Split(os.Getenv("GO_GCFLAGS"), " ") {
		if s == "-N" {
			return 2
		}
	}
	return 1
}