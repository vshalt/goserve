package config

import (
	"fmt"
	"os"
)

type Flags struct {
	Host      string
	Port      string
	Directory string
	Verbose   bool
	SSL       bool
	CertFile  string
	KeyFile   string
}

func Clean(dirs ...string) (string, error) {
	for _, dir := range dirs {
		return dir, nil
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Cannot get cwd: %v", err)
	}
	return cwd, nil
}
