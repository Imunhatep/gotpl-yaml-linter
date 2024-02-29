package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var controlStructureStart = regexp.MustCompile(`^{{-?\s*(if|range|with)\s`)
var controlStructureEnd = regexp.MustCompile(`^{{-?\s*end\s*-?}}`)
var nonControlStructure = regexp.MustCompile(`^{{-?\s*(include|toYaml|nindent)\s`)

func formatLine(line string, indentLevel int) string {
	// Remove leading spaces to reset indentation
	trimmedLine := strings.TrimLeft(line, " ")
	return strings.Repeat("  ", indentLevel) + trimmedLine
}

func isStartControlStructure(line string) bool {
	lineWithoutLeadingSpaces := strings.TrimSpace(line)
	return controlStructureStart.MatchString(lineWithoutLeadingSpaces)
}

func isEndControlStructure(line string) bool {
	lineWithoutLeadingSpaces := strings.TrimSpace(line)
	return controlStructureEnd.MatchString(lineWithoutLeadingSpaces)
}

func isNonControlStructure(line string) bool {
	lineWithoutLeadingSpaces := strings.TrimSpace(line)
	return nonControlStructure.MatchString(lineWithoutLeadingSpaces)
}

func FormatGotpl(data string) (string, error) {
	lines := strings.Split(data, "\n")

	indentLevel := 0
	var formattedLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if isStartControlStructure(trimmed) {
			formattedLines = append(formattedLines, formatLine(line, indentLevel))
			indentLevel++
		} else if isEndControlStructure(trimmed) {
			indentLevel--
			formattedLines = append(formattedLines, formatLine(line, indentLevel))
		} else if isNonControlStructure(trimmed) {
			// Non-control structures and empty lines are indented according to their current block level
			formattedLines = append(formattedLines, formatLine(line, indentLevel))
		} else {
			// Regular lines that are not control structures or non-control structures are treated as text
			formattedLines = append(formattedLines, line)
		}
	}

	return strings.Join(formattedLines, "\n"), nil
}

func FormatFilesInPath(file string, format bool) (bool, error) {
	original, err := os.ReadFile(file)
	if err != nil {
		return false, err
	}

	data, err := FormatGotpl(string(original))
	if err != nil {
		return false, err
	}

	// yaml are invalid
	if string(original) == data {
		fmt.Println("yaml is valid: ", file)
		return true, nil
	}

	// validate, do not change files
	if !format {
		fmt.Println("error! yaml is invalid: ", file)
		return false, nil
	}

	// Write the new content to the file, overwriting existing content
	if err = os.WriteFile(file, []byte(data), 0644); err != nil {
		return false, err
	}

	fmt.Println("yaml linted:", file)

	return true, nil
}
