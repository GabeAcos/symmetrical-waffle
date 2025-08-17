package models

import (
	"database/sql"
	"errors"
)

type Exercise struct {
	ID        int
	Name      string
	Sets      int
	Reps      int
	Weight    int
}

type ExercisesModel struct {
	DB *sql.DB
}

func (m *ExercisesModel) Insert(name string, sets int, reps int, weight int) (int, error) {
	stmt := `INSERT INTO excersises (name, sets, reps, weight)
	VALUES(?, ?, ?, ?)` 

	result, err := m.DB.Exec(stmt, name, sets, reps, weight)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (m *ExercisesModel) Get(id int) (Exercise, error) {
	stmt := `SELECT id, name, sets, reps, weight FROM exercises
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var e Exercise

	err := row.Scan(&e.ID, &e.Name, &e.Sets, &e.Reps, &e.Weight)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Exercise{}, ErrNoRecord
		} else {
			return Exercise{}, err
		}
	}

	return e, nil
}

func (m *ExercisesModel) Latest() ([]Exercise, error) {
	stmt := `SELECT id, name, sets, reps, weight
	FROM exercises
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var excercises []Exercise

	for rows.Next() {
		var e Exercise

		err = rows.Scan(&e.ID, &e.Name, &e.Sets, &e.Reps, &e.Weight)
		if err != nil {
			return nil, err
		}

		excercises = append(excercises, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return excercises, nil
}