//go:build linux

package crossexec_test

import (
	"log"
	"os"
	"path/filepath"

	"go.jcbhmr.com/crossexec"
)

// Use "//go:embed" or something.
var app_sh []byte

// https://pkg.go.dev/go.jcbhmr/crossexec/examples/extract-and-exec
func ExampleExec_extractAndExec() {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	appPath := filepath.Join(userCacheDir, "extract-and-exec-app.sh")

	err = os.WriteFile(appPath, app_sh, 0o755)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(crossexec.Exec(appPath, os.Args, os.Environ()))
}
