package model

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetUp() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "admin",
		AllowNativePasswords: true,
	}
	// dsn := "host=localhost port=3306 user=root password=root dbName=admin sslmode=disabled"
	// "root:root@127.0.0.1:3306/admin?sslmode=disabled"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		_, e := db.Exec(`create table if not exists users (id serial not null unique,name varchar(64) not null,
		email varchar(64) not null unique, password text not null ,primary key(id));`)
		if e != nil {
			panic(e)
		} else {
			fmt.Println("Database initialized")
		}
	}
}
