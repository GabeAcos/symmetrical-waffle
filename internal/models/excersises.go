package models

import (
	"database/sql"
)

type Exercises struct {
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
	return 0, nil
}

func (m *ExercisesModel) Get(id int) (Exercises, error) {
	return Exercises{}, nil
}

func (m *ExercisesModel) Latest() ([]Exercises, error) {
	return nil, nil
}