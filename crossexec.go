//go:build plan9 || unix || windows

package crossexec

// Exec behaves like [syscall.Exec].
func Exec(argv0 string, argv []string, envv []string) error {
	return exec(argv0, argv, envv)
}
