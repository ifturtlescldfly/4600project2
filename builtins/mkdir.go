
package builtins

import (
    "os"
)

// MakeDirectory creates a new directory with the specified name.
func MakeDirectory(args ...string) error {
    for _, path := range args {
        err := os.Mkdir(path, 0755)  // Using a standard permission set
        if err != nil {
            return err
        }
    }
    return nil
}
