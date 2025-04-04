package util

import (
	"path/filepath"
	"strings"
)

func BaseFileNoExt(file string) string {
	return strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
}
