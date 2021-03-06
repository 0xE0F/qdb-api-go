// +build windows

package qdb

import (
	"fmt"
	"path/filepath"
)

func setBinaryRights(path string) {
}

func mkBinaryPath(path string, binary string) string {
	return fmt.Sprintf("%s.exe", filepath.Join(path, binary))
}

func addBinarybase(path *string) {
	*path += "test_qdbd"
}

func addBinaryExtension(path *string) {
	*path += ".exe"
}
