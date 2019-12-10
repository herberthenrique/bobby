package repository

import (
	"fmt"
	"github.com/herberthenrique/bobby/config"
	"github.com/herberthenrique/bobby/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// //TodoRepository will deal with Database operations
// type TodoRepository interface {
// 	GetAllTodos(todo *[]models.Todo) (err error)
// 	CreateATodo(todo *models.Todo) (err error)
// 	GetATodo(todo *models.Todo, id string) (err error)
// 	UpdateATodo(todo *models.Todo, id string) (err error)
// 	DeleteATodo(todo *models.Todo, id string) (err error)
// }

//GetAllTodos will look for all todo
func GetAllTodos(todo *[]models.Todo) (err error) {
	if err := config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//CreateATodo - create a todo
func CreateATodo(todo *models.Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//GetATodo - fetch one todo
func GetATodo(todo *models.Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//UpdateATodo - update a todo
func UpdateATodo(todo *models.Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

//DeleteATodo - delete a todo
func DeleteATodo(todo *models.Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
