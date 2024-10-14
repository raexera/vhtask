package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/raexera/vhtask/internal/application"
	"github.com/raexera/vhtask/internal/infrastructure"
	_interface "github.com/raexera/vhtask/internal/interface"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/raexera/vhtask/docs"
)

// @title VHTask
// @version 1.0
// @description This is a simple To-Do app.

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Connent to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			status VARCHAR(50) CHECK (status IN ('pending', 'in progress', 'completed')) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	repo := infrastructure.NewPostgresTaskRepository(db)
	service := application.NewTaskService(repo)
	handler := _interface.NewTaskHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/tasks", handler.CreateTask)
	e.GET("/tasks", handler.GetAllTasks)
	e.GET("/tasks/:id", handler.GetTaskByID)
	e.PUT("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	} else {
		log.Println("Database connection closed.")
	}
}
