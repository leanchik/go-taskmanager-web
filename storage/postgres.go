package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStorage{DB: db}, nil
}

func (p *PostgresStorage) LoadTasks() ([]Task, error) {
	rows, err := p.DB.Query("SELECT id, title, done, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Title, &t.Done, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (p *PostgresStorage) SaveTasks(tasks []Task) error {
	_, err := p.DB.Exec("DELETE FROM tasks")
	if err != nil {
		return err
	}

	for _, task := range tasks {
		_, err := p.DB.Exec("INSERT INTO tasks (id, title, done, created_at) VALUES ($1, $2, $3, $4)", task.ID, task.Title, task.Done, task.CreatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}
