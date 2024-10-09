package config

import (
	"fmt"
	"os"
	"time"

	"path/filepath"
)

const ReadTimeout = 15 * time.Second
const WriteTimeout = 15 * time.Second
const ReadHeaderTimeout = 30 * time.Second

type Flags struct {
	Host      string
	Port      string
	Directory string
	Verbose   bool
	SSL       bool
	CertFile  string
	KeyFile   string
}

func Clean(dir string) (string, error) {
	absPath, err := filepath.Abs(dir)

	if err != nil {
		return "", fmt.Errorf("Cannot convert to absolute path: %v", err)
	}

	_, err = os.Stat(absPath)
	if err != nil {
		return "", fmt.Errorf("Path %s doesn't exist", dir)
	}
	return absPath, nil
}
