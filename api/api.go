package api

import (
	// "errors"
	"fmt"
	"net/http"

	"github.com/arlisson314/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	// Echo instance
	e := echo.New()
	databaseInit := db.Init()
	database := db.NewStudentHandler(databaseInit)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &API{
		Echo: e,
		DB:   database,
	}
}

func (api *API) StartServer() error {
	// Start server
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	// Routes
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudentbyID)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

// Handler
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve students"})
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}

	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.CreateStudent(student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create student"})
	}
	return c.String(http.StatusCreated, "Student created successfully\n")
}

func (api *API) getStudentbyID(c echo.Context) error {
	id := c.Param("id")
	info := fmt.Sprintf("Get student with ID: %v", id)
	return c.String(http.StatusOK, info)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	info := fmt.Sprintf("Student updated successfully with ID: %v", id)
	return c.String(http.StatusCreated, info)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	info := fmt.Sprintf("Student deleted successfully with ID: %v", id)
	return c.String(http.StatusNoContent, info)
}
