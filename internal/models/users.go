 package models

 import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, name, age, height, weight FROM users
	WHERE id = ?`
	
	row := m.DB.QueryRow(stmt, id)

	var u User

	err := row.Scan(&u.ID, &u.Name, &u.Age, &u.Height, &u.Weight)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrNoRecord
		} else {
			return User{}, err
		}
	}

	return u, nil
 }

 func (m *UserModel) Latest() ([]User, error) {
	stmt := `SELECT id, name, age, height, weight
	FROM users 
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		err = rows.Scan(&u.ID, &u.Name, &u.Age, &u.Height, &u.Weight)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
 }