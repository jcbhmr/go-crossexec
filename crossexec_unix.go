//go:build unix

package crossexec

import (
	"golang.org/x/sys/unix"
)

func crossExec(argv0 string, argv []string, envv []string) error {
	return unix.Exec(argv0, argv, envv)
}
