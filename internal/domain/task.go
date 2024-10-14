package domain

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskRepository interface {
	CreateTask(task *Task) error
	GetTaskByID(id int) (*Task, error)
	GetAllTasks(limit, offset int) ([]*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id int) error
}
