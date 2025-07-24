package handler

import (
	"employee-service-gin/model"
	"employee-service-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var svc = service.NewService()

func AddEmployee(c *gin.Context) {
	var emp model.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	svc.AddEmployee(&emp)
	c.JSON(http.StatusOK, gin.H{"data": "added"})
}

func GetEmployees(c *gin.Context) {
	emps := svc.GetEmployees()
	c.JSON(http.StatusOK, gin.H{"data": emps})
}

func DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	success := svc.DeleteEmployee(id)
	if success {
		c.JSON(http.StatusOK, gin.H{"data": "deleted"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "not found"})
	}
}
