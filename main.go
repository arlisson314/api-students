package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/arlisson314/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudentbyID)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

// Handler
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students\n")
}

func createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	db.AddStudent(student)
	return c.String(http.StatusCreated, "Student created successfully\n")
}

func getStudentbyID(c echo.Context) error {
	id := c.Param("id")
	info := fmt.Sprintf("Get student with ID: %v", id)
	return c.String(http.StatusOK, info)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	info := fmt.Sprintf("Student updated successfully with ID: %v", id)
	return c.String(http.StatusCreated, info)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	info := fmt.Sprintf("Student deleted successfully with ID: %v", id)
	return c.String(http.StatusNoContent, info)
}
