//go:build plan9 || unix || windows

package crossexec

func CrossExec(argv0 string, argv []string, envv []string) error {
	return crossExec(argv0, argv, envv)
}
