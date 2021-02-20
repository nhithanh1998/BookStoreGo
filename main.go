package main

import (
	"bookstore/database"
	"bookstore/resources"
	"bookstore/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	databaseConfig, err := util.LoadDatabaseConfig(".")
	if err != nil {
		fmt.Println("Failed to read config file")
		fmt.Println(err)
		os.Exit(1)
	}
	if err := database.OpenConnectionToDatabase(&databaseConfig); err != nil {
		fmt.Println("Failed to connect to database")
		os.Exit(1)
	}
	database.MigrateModel()
	router := gin.Default()
	resources.InitCategories(router)
	router.Run(":3000")
}
