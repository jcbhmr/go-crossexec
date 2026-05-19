//go:build plan9 || unix || windows

package crossexec_test

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

// https://pkg.go.dev/go.jcbhmr/crossexec/examples/go-version
func ExampleExec_goVersion() {

	log.Fatal(crossexec.Exec(MustLookPath("go"), []string{"go", "version"}, os.Environ()))
}
