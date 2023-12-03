
package builtins

import (
    "io/ioutil"
    "os"
    "testing"
)

func TestRemoveFileOrDirectory(t *testing.T) {
    // Create a temporary file to test removal
    tempFile, err := ioutil.TempFile("", "test")
    if err != nil {
        t.Fatalf("Unable to create temporary file: %v", err)
    }
    tempFilePath := tempFile.Name()
    tempFile.Close()

    // Remove the temporary file
    err = RemoveFileOrDirectory(tempFilePath)
    if err != nil {
        t.Errorf("Failed to remove file: %v", err)
    }

    // Check if the file still exists
    if _, err := os.Stat(tempFilePath); !os.IsNotExist(err) {
        t.Errorf("File %v still exists after removal", tempFilePath)
    }
}
