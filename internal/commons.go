package internal

import (
	"os"
	"strings"
)

func ReadFileAsLines(fname string) ([]string, error) {
	inputBytes, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var out []string
	lines := strings.Split(string(inputBytes), "\n")
	for _, line := range lines {
		if line != "" {
			out = append(out, line)
		}
	}

	return out, nil
}

func ReadFileAsString(fname string) (string, error) {
	inputBytes, err := os.ReadFile(fname)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(inputBytes)), nil
}
