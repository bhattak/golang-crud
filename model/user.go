package model

import (
	"errors"
	"go-db/database"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	statement := `insert into users (name,email,password) values(?,?,?);`
	_, err := database.Db.Exec(statement, user.Name, user.Email, user.Password)
	return err
}

func GetUser(id int) (User, error) {
	var user User
	statement := `select * from users where id = ?;`
	rows, err := database.Db.Query(statement, id)

	if err != nil {
		return User{}, err
	}

	if !rows.Next() {
		return User{}, errors.New("User not found")
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func GetAllUser() ([]User, error) {
	var userList []User
	statement := `select * from users;`
	rows, err := database.Db.Query(statement)
	if err != nil {
		return []User{}, err
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return []User{}, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}

func UpdateUser(id int, pass string) error {
	statement := `update users set password = ? where id=?;`
	_, err := database.Db.Query(statement, pass, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	statement := `delete from users where id=?;`
	_, err := database.Db.Query(statement, id)
	if err != nil {
		return err
	}
	return nil
}
