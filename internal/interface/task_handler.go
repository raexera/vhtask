package _interface

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/raexera/vhtask/internal/application"
	"github.com/raexera/vhtask/internal/domain"
)

type TaskHandler struct {
	service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask creates a new task.
// @Summary Create a new task
// @Description Create a new task with title, description, and status
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body domain.Task true "Task Info"
// @Success 201 {object} domain.Task
// @Failure 400 {object} map[string]string "Invalid payload"
// @Failure 500 {object} map[string]string "Failed to create task"
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c echo.Context) error {
	task := new(domain.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}

	task.Status = "pending"
	if err := h.service.CreateTask(task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

// GetTaskByID gets a task by ID.
// @Summary Get task by ID
// @Description Get a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} domain.Task
// @Failure 404 {object} map[string]string "Task not found"
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	return c.JSON(http.StatusOK, task)
}

// GetAllTasks gets all tasks.
// @Summary Get all tasks
// @Description Get all tasks with pagination
// @Tags tasks
// @Accept json
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} domain.Task
// @Failure 500 {object} map[string]string "Failed to retrieve tasks"
// @Router /tasks [get]
func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	tasks, err := h.service.GetAllTasks(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to retrieve tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

// UpdateTask updates an existing task.
// @Summary Update an existing task
// @Description Update task title, description, or status
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body domain.Task true "Updated Task Info"
// @Success 200 {object} domain.Task
// @Failure 400 {object} map[string]string "Invalid payload"
// @Failure 500 {object} map[string]string "Failed to update task"
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	task := new(domain.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	task.ID = id

	if err := h.service.UpdateTask(task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update task"})
	}

	return c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task by ID.
// @Summary Delete a task by ID
// @Description Delete a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 204 "Task deleted"
// @Failure 500 {object} map[string]string "Failed to delete task"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
	}

	return c.JSON(http.StatusNoContent, nil)
}
