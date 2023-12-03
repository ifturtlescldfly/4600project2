
package builtins

import (
    "bytes"
    "io/ioutil"
    "os"
    "strings"
    "testing"
)

func TestListDirectory(t *testing.T) {
    // Create a temporary directory and file for testing
    tempDir, err := ioutil.TempDir("", "testLs")
    if err != nil {
        t.Fatalf("Unable to create temporary directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    tempFile, err := ioutil.TempFile(tempDir, "testfile")
    if err != nil {
        t.Fatalf("Unable to create temporary file: %v", err)
    }
    tempFileName := tempFile.Name()
    tempFile.Close()

    // Testing the ListDirectory function
    var b bytes.Buffer
    err = ListDirectory(&b, tempDir)
    if err != nil {
        t.Fatalf("ListDirectory returned an error: %v", err)
    }

    // Check if the temporary file is listed
    output := b.String()
    if !strings.Contains(output, tempFileName) {
        t.Errorf("ListDirectory did not list the expected file. Got: %s", output)
    }
}
