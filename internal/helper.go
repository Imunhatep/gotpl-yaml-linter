package internal

import (
	"errors"
	"fmt"
	"path/filepath"
)

func ListFilesInDir(path, pattern *string) ([]string, error) {
	// Get absolute path
	absPath, err := filepath.Abs(*path)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Error getting absolute path: %v\n", err))
	}

	// Find files matching pattern in the directory
	matches, err := filepath.Glob(filepath.Join(absPath, *pattern))
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Error matching pattern: %v\n", err))
	}

	return matches, nil
}
