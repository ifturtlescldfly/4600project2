
package builtins

import (
    "os"
    "path/filepath"
    "testing"
)

func TestMakeDirectory(t *testing.T) {
    // Creating a temporary directory path for testing
    tempDirPath := filepath.Join(os.TempDir(), "testMkdir")

    // Removing the temporary directory at the end of the test
    defer os.RemoveAll(tempDirPath)

    // Testing the MakeDirectory function
    err := MakeDirectory(tempDirPath)
    if err != nil {
        t.Fatalf("Failed to create directory: %v", err)
    }

    // Check if the directory was created
    if _, err := os.Stat(tempDirPath); os.IsNotExist(err) {
        t.Errorf("Directory was not created: %s", tempDirPath)
    }
}
