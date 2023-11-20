package database

import (
	"fmt"
	"log"
	"math/rand"
)

var db = InitDB()

func InsertTodo(returningId *string, title string) error {
	query := `
	INSERT INTO todos (id, title, done) 
	VALUES ($1, $2, $3)
	RETURNING id`
	id := fmt.Sprint(rand.Int())
	err := db.QueryRow(query, id, title, false).Scan(returningId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()
	return nil
}

func UpdateTodo(id string, newTitle string) int64 {
	query := `
	UPDATE todos
	SET title = $2
	WHERE id = $1
	RETURNING id;`
	res, err := db.Exec(query, id, newTitle)
	if err != nil {
		fmt.Println("No rows were affected")
		fmt.Println(err)
		return 0
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("No rows were affected")
		fmt.Println(err)
		return 0
	}

	log.Print("Successfully updated", id)
	return rowsAffected
}

func DeleteTodo(id string) int64 {
	query := `
	DELETE FROM todos
	WHERE id = $1
	RETURNING id`
	res, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("No rows were affected")
		fmt.Println(err)
		return 0
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("No rows were affected")
		fmt.Println(err)
		return 0
	}
	log.Print("Successfully deleted", id)
	return rowsAffected
}
