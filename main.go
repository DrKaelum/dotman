package main

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Get home dir path
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// open input file
	fi, err := os.Open(filepath.Join(home, ".zshrc"))
	if err != nil {
		panic(err)
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// Create backups file if it doesn't exist
	if err := os.MkdirAll(filepath.Join(home, "backups"), 0o755); err != nil {
		panic(err)
	}

	timestamp := time.Now().Format("20060102-150405")
	backupPath := filepath.Join(home, "backups", ".zshrc."+timestamp+".backup")

	// open output file
	fo, err := os.OpenFile(backupPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := io.Copy(fo, fi); err != nil {
		panic(err)
	}
}
