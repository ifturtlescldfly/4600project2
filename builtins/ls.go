
package builtins

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// ListDirectory lists the files and directories in the current or specified directory.
func ListDirectory(w io.Writer, args ...string) error {
    var dir string
    if len(args) > 0 {
        dir = args[0]
    } else {
        var err error
        dir, err = os.Getwd()
        if err != nil {
            return err
        }
    }

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return err
    }

    for _, file := range files {
        _, err = fmt.Fprintln(w, file.Name())
        if err != nil {
            return err
        }
    }

    return nil
}
