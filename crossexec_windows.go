package crossexec

import (
	"os"
	"runtime"
	"sync"

	"go.jcbhmr.com/crossexec/internal/consoleapi"
	"golang.org/x/sys/windows"
)

const windowsTRUE = 1

var getIgnoreCtrlHandler = sync.OnceValue(func() consoleapi.PHANDLER_ROUTINE {
	return consoleapi.PHANDLER_ROUTINE(windows.NewCallback(func(_ uint32) uintptr {
		return windowsTRUE
	}))
})

var crossExecMu sync.Mutex

func crossExec(argv0 string, argv []string, envv []string) (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	crossExecMu.Lock()
	defer crossExecMu.Unlock()

	err = consoleapi.SetConsoleCtrlHandler(getIgnoreCtrlHandler(), true)
	if err != nil {
		return err
	}

	process, err := os.StartProcess(argv0, argv, &os.ProcAttr{
		Env:   envv,
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})
	if err != nil {
		return err
	}

	state, err := process.Wait()
	if err != nil {
		panic(err)
	}

	os.Exit(state.ExitCode())
	panic("noreturn")
}
