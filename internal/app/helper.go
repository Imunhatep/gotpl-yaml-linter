package app

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"path/filepath"
)

func ListFilesInDir(path, pattern string) ([]string, error) {
	// Get absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Error getting absolute path: %v\n", err))
	}

	// Find files matching pattern in the directory
	matches, err := filepath.Glob(filepath.Join(absPath, pattern))
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Error matching pattern: %v\n", err))
	}

	return matches, nil
}

func ToYaml(data interface{}) string {
	b, err := yaml.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("failed parsing to yaml")
		return ""
	}

	return string(b)
}
