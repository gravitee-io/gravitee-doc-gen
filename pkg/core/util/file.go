package util

import (
	"os"
	"path/filepath"
	"strings"
)

func BaseFileNoExt(file string) string {
	return strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}
