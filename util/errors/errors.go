// This is pretty much the standard go errors package, but includes a
// stack trace in the error message too.
package errors

import (
	stderr "errors"
	"fmt"
	"runtime/debug"
)

func New(data string) error {
	return stderr.New(fmt.Sprintf("%s\n%s", data, string(debug.Stack())))
}
