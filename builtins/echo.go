
package builtins

import (
    "fmt"
    "io"
)

// Echo prints the arguments to the writer.
func Echo(w io.Writer, args ...string) error {
    _, err := fmt.Fprintln(w, args)
    return err
}
