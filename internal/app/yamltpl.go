package app

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"regexp"
	"strings"
)

var controlStructureStart = regexp.MustCompile(`^{{-?\s*(if|range|with|define)\s`)
var controlStructureContinue = regexp.MustCompile(`^{{-?\s*(else)\s`)
var controlStructureEnd = regexp.MustCompile(`^{{-?\s*end\s*-?}}`)
var nonControlStructure = regexp.MustCompile(`^{{-?\s*(include|toYaml|nindent)\s`)

// FormatYamlTpl formats a yaml template string
func FormatYamlTpl(yamlTpl string) (string, error) {
	lines := strings.Split(yamlTpl, "\n")

	indentLevel := 0
	var formattedLines []string
	for i, line := range lines {
		trimmed := strings.TrimSpace(strings.Replace(line, "\t", "\n", -1))
		if isStartControlStructure(trimmed) {
			formattedLines = append(formattedLines, formatLine(line, indentLevel))
			indentLevel++
		} else if isContinueControlStructure(trimmed) {
			formattedLines = append(formattedLines, formatLine(line, indentLevel-1))
		} else if isNonControlStructure(trimmed) {
			// Non-control structures and empty lines are indented according to their current block level
			formattedLines = append(formattedLines, formatLine(line, indentLevel))
		} else if isEndControlStructure(trimmed) {
			// End control structures are indented according to their current block level
			indentLevel--

			if indentLevel < 0 {
				indentLevel = 0
				log.Warn().Msgf("Seems closing structure has no opening. Invalid gotpl structure at line ~%d: %s", i, line)
			}

			formattedLines = append(formattedLines, formatLine(line, indentLevel))
		} else {
			// Regular lines that are not control structures or non-control structures are treated as text
			formattedLines = append(formattedLines, line)
		}
	}

	return strings.Join(formattedLines, "\n"), nil
}

// FormatYamlTplFile formats a yaml file
func FormatYamlTplFile(file string, format, output bool) (bool, error) {
	original, err := os.ReadFile(file)
	if err != nil {
		return false, err
	}

	data, err := FormatYamlTpl(string(original))
	if err != nil {
		return false, err
	}

	// output expected file formatting
	if output {
		fmt.Printf("\nexpected yaml [%s] tpl formtting:\n%s\n\n", file, data)
	}

	// yaml are invalid
	if string(original) == data {
		log.Info().Str("file", file).Msgf("yaml template is valid")
		return true, nil
	}

	// validate, do not change files
	if !format {
		log.Error().Str("file", file).Msgf("error! yaml is invalid")
		return false, nil
	}

	// Write the new content to the file, overwriting existing content
	if err = os.WriteFile(file, []byte(data), 0644); err != nil {
		return false, err
	}

	log.Info().Str("file", file).Msg("linted")

	return true, nil
}

func formatLine(line string, indentLevel int) string {
	// Remove leading spaces to reset indentation
	trimmedLine := strings.TrimLeft(line, " ")
	return strings.Repeat("  ", indentLevel) + trimmedLine
}

func isStartControlStructure(line string) bool {
	lineWithoutLeadingSpaces := strings.TrimSpace(line)
	return controlStructureStart.MatchString(lineWithoutLeadingSpaces)
}

func isContinueControlStructure(line string) bool {
	lineWithoutLeadingSpaces := strings.TrimSpace(line)
	return controlStructureContinue.MatchString(lineWithoutLeadingSpaces)
}

func isEndControlStructure(line string) bool {
	lineWithoutLeadingSpaces := strings.TrimSpace(line)
	return controlStructureEnd.MatchString(lineWithoutLeadingSpaces)
}

func isNonControlStructure(line string) bool {
	lineWithoutLeadingSpaces := strings.TrimSpace(line)
	return nonControlStructure.MatchString(lineWithoutLeadingSpaces)
}
