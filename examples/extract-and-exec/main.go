//go:build unix

package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"

	"go.jcbhmr.com/crossexec"
)

//go:embed app.sh
var app_sh []byte

func main() {
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
