
package builtins

import (
    "bytes"
    "testing"
)

func TestEcho(t *testing.T) {
    var b bytes.Buffer
    args := []string{"hello", "world"}

    err := Echo(&b, args...)
    if err != nil {
        t.Errorf("Echo returned an error: %v", err)
    }

    got := b.String()
    want := "hello world\n"
    if got != want {
        t.Errorf("Echo = %q, want %q", got, want)
    }
}
