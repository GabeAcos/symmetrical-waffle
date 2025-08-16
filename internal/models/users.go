 package models

 import (
	"database/sql"
 )

 type User struct {
	ID		int
	Name    string
	Age     int
	Height	int
	Weight	int
 }



 type UserModel struct {
	DB *sql.DB
 }

 func (m *UserModel) Insert(name string, age int, height int, weight int) (int, error) {
	stmt := `INSERT INTO users (name, age, height, weight)
	VALUES(?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, name, age, height, weight)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
 }

 func (m *UserModel) Get(id int) (User, error) {
	return User{}, nil
 }

 func (m *UserModel) Latest() ([]User, error) {
	return nil, nil
 }