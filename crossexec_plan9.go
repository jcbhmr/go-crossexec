package crossexec

import (
	// golang.org/x/sys/plan9 doesn't have Exec. It should. syscall does have Exec.
	"syscall"
)

func exec(argv0 string, argv []string, envv []string) error {
	return syscall.Exec(argv0, argv, envv)
}
