package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herberthenrique/bobby/models"
	"github.com/herberthenrique/bobby/repository"
)

//TodoController Main controller of Todo resource
type TodoController struct{}

//Show - List all todos
func (con TodoController) Show(c *gin.Context) {
	var todo []models.Todo
	err := repository.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Create - Create a Todo
func (con TodoController) Create(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := repository.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Get - Get a particular Todo with id
func (con TodoController) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := repository.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Update -  Update an existing Todo
func (con TodoController) Update(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := repository.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = repository.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//Delete -  Delete a Todo
func (con TodoController) Delete(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := repository.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
