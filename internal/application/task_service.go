package application

import "github.com/raexera/vhtask/internal/domain"

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *domain.Task) error {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTaskByID(id int) (*domain.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) GetAllTasks(limit, offset int) ([]*domain.Task, error) {
	return s.repo.GetAllTasks(limit, offset)
}

func (s *TaskService) UpdateTask(task *domain.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}
