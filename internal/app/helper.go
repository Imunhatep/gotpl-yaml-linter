package app

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

// ListFilesInDir lists all files in a directory matching a pattern
func ListFilesInDir(path, pattern string) ([]string, error) {
	// Get absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return []string{}, fmt.Errorf("error getting absolute path: %v", err)
	}

	// Find files matching pattern in the directory
	matches, err := filepath.Glob(filepath.Join(absPath, pattern))
	if err != nil {
		return []string{}, fmt.Errorf("error matching pattern: %v", err)
	}

	return matches, nil
}

// ToYaml converts data to yaml
func ToYaml(data interface{}) string {
	b, err := yaml.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("failed parsing to yaml")
		return ""
	}

	return string(b)
}
