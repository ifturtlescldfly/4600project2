
package builtins

import (
    "bytes"
    "os"
    "testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
    var b bytes.Buffer

    err := PrintWorkingDirectory(&b)
    if err != nil {
        t.Errorf("PrintWorkingDirectory returned an error: %v", err)
    }

    got := b.String()
    want := os.Getenv("PWD") + "\n"
    if got != want {
        t.Errorf("PrintWorkingDirectory = %q, want %q", got, want)
    }
}
