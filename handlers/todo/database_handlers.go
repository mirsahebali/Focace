package todo

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"mirsahebali/focace/database"
)

var db = database.InitDB()

func ReadTodos(email string) []Todo {
	var itemList []Todo
	sqlStatement := `SELECT * FROM todos WHERE email = $1`
	rows, err := db.Query(sqlStatement, email)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done, &todo.Email, &todo.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		itemList = append(itemList, todo)
	}
	return itemList
}

func InsertTodo(todo *Todo, title, email string) error {
	query := `
	INSERT INTO todos (id, title, done, createdat,email) 
	VALUES ($1, $2, $3, $4,$5)
	RETURNING id,title,done,email`
	id := fmt.Sprint(rand.Int())
	err := db.QueryRow(query, id, title, false, time.Now(), email).Scan(&todo.Id, &todo.Title, &todo.Done, &todo.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (t *Todo) UpdateTodo(id string, newTitle string) int64 {
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

func (t *Todo) DeleteTodo(id string) int64 {
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
