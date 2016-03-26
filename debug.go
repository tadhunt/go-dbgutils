package debug

import (
	"fmt"
	"runtime"
)

func FuncName(skip int, verbose bool) string {
	pc, file, line, ok := runtime.Caller(skip + 1)

	if ok {
		if verbose {
			return fmt.Sprintf("[%s:%d]%s", file, line, runtime.FuncForPC(pc).Name())
		}

		return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	}

	return "<unknown>"
}
