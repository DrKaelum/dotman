package main

import (
	"io"
	"os"
	"path/filepath"
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

	// open output file
	fo, err := os.Create(filepath.Join(home, "backups/.zshrc.backup"))
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}
