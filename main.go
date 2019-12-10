package main

import (
	"fmt"
	"github.com/herberthenrique/boilerplate_gin/config"
	"github.com/herberthenrique/boilerplate_gin/models"
	"github.com/herberthenrique/boilerplate_gin/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {

	// Creating a connection to the database
	config.DB, err = gorm.Open("postgres", config.DabataseURL(config.BuildDatabaseConfiguration()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run()
}
