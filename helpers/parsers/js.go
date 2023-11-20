package parsers

import (
	"html/template"
	"os"
)

func JSParser(path string) (template.JS, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return template.JS(string(file)), nil
}
