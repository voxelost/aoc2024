package internal

import (
	"os"
	"strings"
)

func ReadFileAsString(fname string) (string, error) {
	inputBytes, err := os.ReadFile(fname)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(inputBytes)), nil
}
