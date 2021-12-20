package router

import (
	c "github.com/KeshikaGupta20/Postgresql_GO/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.POST("/insertpro", c.AddEmployee)

	router.DELETE("/deletepro/:id", c.DeleteEmployee)

	router.GET("/getpro", c.GetEmploy)

	router.GET("/updatepro/:id", c.GetEmployees)
}
