package parsers

import (
	"fmt"
	"os"
	"strings"
)

const MODNAME = "focace"

func GetEnvVars(key string) string {
	envValue := os.Getenv(key)
	if envValue != "" {
		return envValue
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	dirSlice := strings.Split(wd, "/")
	var modIndex int
	for i, dir := range dirSlice {
		if MODNAME == dir {
			modIndex = i
			break
		}
	}
	envFileIndex := modIndex + 1
	rootFolder := dirSlice[:envFileIndex]
	rootPath := strings.Join(rootFolder, "/")
	envFile, err := os.ReadFile(rootPath + "/.env")
	if err != nil {
		fmt.Println("env file not found")
		return ""
	}
	envLines := strings.Split(string(envFile), "\n")

	var envKeyValue string

	for _, envKeyVals := range envLines {
		if strings.Contains(envKeyVals, key) {
			envKeyValue = envKeyVals
		}
	}
	values := strings.Split(envKeyValue, "=")[1:]

	return strings.Join(values, "=")
}
