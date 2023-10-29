package routes

import (
	"notepad-api/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/new-note", controllers.AddNewNote)
	e.GET("/notes", controllers.ReadAllNote)
	e.GET("/note/:id", controllers.ReadOneNote)
	e.DELETE("/delete-note/:id", controllers.DeleteOneNote)

	e.POST("/create-user", controllers.CreateUser)
	e.GET("/user/:id", controllers.ProfileUser)
	e.DELETE("/delete-user/:id", controllers.DeleteOneUser)
	e.POST("/login", controllers.Login)
}
