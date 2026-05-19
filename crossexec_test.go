//go:build plan9 || unix || windows

package crossexec_test

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"go.jcbhmr.com/crossexec"
)

func MustLookPath(file string) string {
	p, err := exec.LookPath(file)
	if err != nil {
		panic(err)
	}
	return p
}

func ExampleExec_goVersion() {

	log.Fatal(crossexec.Exec(MustLookPath("go"), []string{"go", "version"}, os.Environ()))
}

// Use "//go:embed" or something.
var app_sh []byte

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
