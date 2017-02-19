package qsync

import (
	"fmt"
	"os"
	"path/filepath"
)

func FileExist(file string) (string, bool) {
	path, err := filepath.Abs(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "qsync:%v\n", err)
		os.Exit(1)
	}
	if _, err := os.Stat(path); err != nil {
		fmt.Println(err)
		return path, false
	}
	return path, true
}
