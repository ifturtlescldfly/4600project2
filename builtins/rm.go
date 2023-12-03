
package builtins

import (
    "os"
)

// RemoveFileOrDirectory removes the specified file or directory.
func RemoveFileOrDirectory(args ...string) error {
    for _, path := range args {
        err := os.Remove(path)
        if err != nil {
            return err
        }
    }
    return nil
}
