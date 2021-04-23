package database

import (
	"database/sql"
	"fmt"
	"time"
)

func Start() {
	db, err := sql.Open("mysql", "root:@protocol(localhost)/dizzle?param=value")

	if err != nil {
		fmt.Println(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
