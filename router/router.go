package router

import (
	"assigment_project_rest_api/controller"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	// app := gin.Default()
	routes := app.RouterGroup
	v1 := routes.Group("/student")
	v2 := routes.Group("/students")
	v1.POST("/", controller.CreateStudent)
	v1.DELETE("/:id", controller.DeleteStudentByID)
	v1.PUT("/:id", controller.UpdateStudentByID)

	v2.GET("/", controller.GetAllStudent)

}
