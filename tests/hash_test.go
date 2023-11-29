package tests

import (
	"testing"

	"mirsahebali/focace/helpers/encrypt"
)

func TestPassword(t *testing.T) {
	encrypt.HashPassword("foo+bar+bazz", "My_Secret")
}
