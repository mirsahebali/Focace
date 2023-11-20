package tests

import (
	"fmt"
	"mirsahebali/focace/helpers/parsers"
	"testing"
)

func TestEnv(t *testing.T) {
	want := parsers.GetEnvVars("FOCACE_AUTH_SECRET")

	fmt.Println(want)
}
