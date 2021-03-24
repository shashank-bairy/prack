package utils

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectPath = filepath.Join(filepath.Dir(b), "..")
)
