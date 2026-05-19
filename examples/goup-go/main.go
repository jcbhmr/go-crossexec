/*
TODO

		$ go version
		go version go1.26.3 linux/amd64
	    $ go run go.jcbhmr.com/crossexec/examples/goup-go@latest +go1.26.0 version
		go version go1.26.0 linux/amd64
*/
package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"go.jcbhmr.com/crossexec"
)

func main() {
	var goPath string
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "+") {
		err := os.Setenv("GOTOOLCHAIN", strings.TrimPrefix(os.Args[1], "+"))
		if err != nil {
			log.Fatal(err)
		}
	}
	var err error
	goPath, err = exec.LookPath("go")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(crossexec.Exec(goPath, os.Args, os.Environ()))
}
