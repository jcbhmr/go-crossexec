/*
TODO

		$ go version
		go version go1.26.3 linux/amd64
	    $ go run go.jcbhmr.com/crossexec/examples/go-version@latest
		go version go1.26.3 linux/amd64
*/
package main

import (
	"log"
	"os"
	"os/exec"

	"go.jcbhmr.com/crossexec"
)

func MustLookPath(file string) string {
	p, err := exec.LookPath(file)
	if err != nil {
		panic(err)
	}
	return p
}

func main() {
	log.Fatal(crossexec.Exec(MustLookPath("go"), []string{"go", "version"}, os.Environ()))
}
