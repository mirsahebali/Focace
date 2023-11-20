package auth

import (
	"fmt"
	"math/rand"
	"mirsahebali/focace/database"
	"time"
)

type User struct {
	Id           string `json:"id"`
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Passwordhash string `json:"passwordhash"`
	CreatedAt    string `json:"createdAt"`
}

var userList []User
var db = database.InitDB()

func fetchUsers() {

	rows, err := db.Query(`SELECT id,email,passwordhash,createdat,fullname,username FROM users;`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email, &user.Passwordhash, &user.CreatedAt, &user.Fullname, &user.Username)
		userList = append(userList, user)
	}
}

func GetUserObject(email string) (User, bool) {
	fetchUsers()
	for _, user := range userList {
		if user.Email == email {
			fmt.Println("user found")
			return user, true
		}
	}

	return User{}, false
}

func (u *User) ValidatePasswordHash(passwordHash string) bool {
	return u.Passwordhash == passwordHash
}

func AddUserObject(u User) bool {

	for _, user := range userList {
		if u.Email == user.Email {
			fmt.Println("User already exists")
			return false
		}
	}
	sqlStatement := `INSERT INTO users (id,  email, username, fullname, createdat,passwordhash) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	id := fmt.Sprint(rand.Int31())
	err := db.QueryRow(sqlStatement, id, u.Email, u.Username, u.Fullname, time.Now(), u.Passwordhash).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	userList = append(userList, u)
	return true
}
