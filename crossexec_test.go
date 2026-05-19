//go:build plan9 || unix || windows

package crossexec_test

import (
	"go.jcbhmr.com/crossexec"
)

// [./examples/go-version]
//
// [./examples/go-version]: ./examples/go-version
func ExampleExec_goVersion() {
	_ = crossexec.Exec
}

// [./examples/extract-and-exec]
//
// [./examples/extract-and-exec]: ./examples/extract-and-exec
func ExampleExec_extractAndExec() {
	_ = crossexec.Exec
}
