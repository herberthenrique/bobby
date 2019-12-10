package routes

import (
	"github.com/gin-gonic/gin"
)

//SetupRouter will setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")

	SetupTodoRouter(v1)

	return r
}
