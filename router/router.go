package router

import (
	"employee-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	employee := r.Group("/employees")
	{
		employee.POST("/", handler.AddEmployee)
		employee.GET("/", handler.GetEmployees)
		employee.DELETE("/:id", handler.DeleteEmployee)
	}

	return r
}
