package tests

import (
	"fmt"
	"testing"

	"mirsahebali/focace/database"
	"mirsahebali/focace/helpers/parsers"

	_ "github.com/lib/pq"
)

var (
	user = parsers.GetEnvVars("DB_USER")
	pass = parsers.GetEnvVars("DB_PASS")
	host = parsers.GetEnvVars("DB_HOST")
	port = parsers.GetEnvVars("DB_PORT")
	name = parsers.GetEnvVars("DB_NAME")
)

func TestQuery(t *testing.T) {
}

func TestMutation(t *testing.T) {
	var id string
	err := database.InsertTodo(&id, "newtoodo 88")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)
}
