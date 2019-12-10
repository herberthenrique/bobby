package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herberthenrique/boilerplate_gin/controllers"
)

//SetupTodoRouter will setup router
func SetupTodoRouter(e *gin.RouterGroup) {
	todo := e.Group("todo")
	{
		todo.GET("/", controllers.GetTodos)
		todo.POST("/", controllers.CreateATodo)
		todo.GET("/:id", controllers.GetATodo)
		todo.PUT("/:id", controllers.UpdateATodo)
		todo.DELETE("/:id", controllers.DeleteATodo)
	}

}
