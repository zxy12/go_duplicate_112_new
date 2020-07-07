// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"cmd/asm/internal/arch"
	"cmd/asm/internal/flags"
	"cmd/asm/internal/lex"
	"cmd/internal/bio"
	"cmd/internal/obj"
	"cmd/internal/objabi"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("asm: ")

	GOARCH := objabi.GOARCH

	//println("GOARCH=", GOARCH)

	architecture := arch.Set(GOARCH)
	if architecture == nil {
		log.Fatalf("unrecognized architecture %s", GOARCH)
	}
	flags.Parse()

	ctxt := obj.Linknew(architecture.LinkArch)
	if *flags.PrintOut {
		ctxt.Debugasm = 1
	}

	ctxt.Flag_dynlink = *flags.Dynlink
	ctxt.Flag_shared = *flags.Shared || *flags.Dynlink
	ctxt.Bso = bufio.NewWriter(os.Stdout)
	defer ctxt.Bso.Flush()

	architecture.Init(ctxt)

	// Create object file, write header.
	out, err := os.Create(*flags.OutputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer bio.MustClose(out)
	buf := bufio.NewWriter(bio.MustWriter(out))

	// 输出一行字到x.o文件
	if !*flags.SymABIs {
		fmt.Fprintf(buf, "go object %s %s %s, modified by zxy\n", objabi.GOOS, objabi.GOARCH, objabi.Version)
		fmt.Fprintf(buf, "!\n")
	}

	var ok, diag bool
	var failedFile string

	for _, f := range flag.Args() {
		lexer := lex.NewLexer(f)
		_ = lexer
		//parser := asm.NewParser(ctxt, architecture, lexer)
		//_, _ = lexer, parser

	}
	/*


			ctxt.DiagFunc = func(format string, args ...interface{}) {
				diag = true
				log.Printf(format, args...)
			}
			if *flags.SymABIs {
				ok = parser.ParseSymABIs(buf)
			} else {
				pList := new(obj.Plist)
				pList.Firstpc, ok = parser.Parse()
				// reports errors to parser.Errorf
				if ok {
					obj.Flushplist(ctxt, pList, nil, "")
				}
			}
			if !ok {
				failedFile = f
				break
			}
		}
		if ok && !*flags.SymABIs {
			obj.WriteObjFile(ctxt, buf)
		}
		if !ok || diag {
			if failedFile != "" {
				log.Printf("assembly of %s failed", failedFile)
			} else {
				log.Print("assembly failed")
			}
			out.Close()
			os.Remove(*flags.OutputFile)
			os.Exit(1)
		}

	*/
	_, _, _ = ok, diag, failedFile
	buf.Flush()
}
