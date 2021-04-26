package database

import (
	"database/sql"
	"fmt"
	"time"
)

var connection *sql.DB

func Start() {
	db, err := sql.Open("mysql", "root:@protocol(localhost)/dizzle?param=value")

	connection = db

	if err != nil {
		fmt.Println(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err != nil {
		fmt.Println(err.Error())
	}
}

type UserStruct struct {
	name   string
	uid    string
	reason string
}

func createBan(user *UserStruct) {
	rows, err := connection.Query("INSERT INTO bans (name, uid, reason) VALUES (?, ?, ?)", user.name, user.reason, user.uid)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(rows)
}
