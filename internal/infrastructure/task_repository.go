package infrastructure

import (
	"database/sql"

	"github.com/raexera/vhtask/internal/domain"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) domain.TaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) CreateTask(task *domain.Task) error {
	query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID)
}

func (r *PostgresTaskRepository) GetTaskByID(id int) (*domain.Task, error) {
	task := &domain.Task{}
	query := "SELECT id, title, description, status FROM tasks WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	return task, err
}

func (r *PostgresTaskRepository) GetAllTasks(limit, offset int) ([]*domain.Task, error) {
	query := "SELECT id, title, description, status FROM tasks LIMIT $1 OFFSET $2"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*domain.Task{}
	for rows.Next() {
		task := &domain.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *PostgresTaskRepository) UpdateTask(task *domain.Task) error {
	query := "UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4"
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
	return err
}

func (r *PostgresTaskRepository) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
