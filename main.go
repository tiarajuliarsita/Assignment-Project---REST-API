package main

import (
	"assigment_project_rest_api/database"
	"assigment_project_rest_api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	app := gin.Default()
	router.Routes(app)

	app.Run(":8080")
}
