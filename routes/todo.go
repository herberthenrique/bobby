package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herberthenrique/bobby/controllers"
)

//SetupTodoRouter will setup router
func SetupTodoRouter(e *gin.RouterGroup) {
	TodoController := controllers.TodoController{}

	todo := e.Group("todo")
	{
		todo.GET("/", TodoController.Show)
		todo.POST("/", TodoController.Create)
		todo.GET("/:id", TodoController.Get)
		todo.PUT("/:id", TodoController.Update)
		todo.DELETE("/:id", TodoController.Delete)
	}

}
