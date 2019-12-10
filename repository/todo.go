package repository

import (
	"fmt"
	"github.com/herberthenrique/bobby/config"
	"github.com/herberthenrique/bobby/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

// TodoRepository will deal with Database operations
type TodoRepository struct {
	DB *gorm.DB
}

func (repo TodoRepository) init() {
	repo.DB = config.DB
}

//List will look for all todo
func (repo TodoRepository) List(todo *[]models.Todo) (err error) {
	if err := config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//Create - create a todo
func (repo TodoRepository) Create(todo *models.Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//Get - fetch one todo
func (repo TodoRepository) Get(todo *models.Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//Update - update a todo
func (repo TodoRepository) Update(todo *models.Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

//Delete - delete a todo
func (repo TodoRepository) Delete(todo *models.Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
