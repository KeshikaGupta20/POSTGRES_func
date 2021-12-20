package controller

import (
	"net/http"

	"github.com/KeshikaGupta20/Postgresql_GO/database"
	"github.com/KeshikaGupta20/Postgresql_GO/models"
	"github.com/gin-gonic/gin"

	
)

func AddEmployee(c *gin.Context) {

	db := database.DB

	emp := models.Employee{}

	c.BindJSON(&emp)

	db.Exec("Call Employ($1, $2) ", "name", "empid").Scan(&models.Employee{})

	c.JSON(http.StatusOK, gin.H{"message": "Employee table created successfully"})

}

func DeleteEmployee(c *gin.Context){

	ID := c.Param("empid")

	db := database.DB

	var employee models.Employee

	db.Find(&employee, ID)

	db.Delete(&employee, ID)

	c.JSON(http.StatusOK, gin.H{"message": "Employee table created successfully"})

}



func GetEmployees(c *gin.Context) {

	db := database.DB

	var employee models.Employee

	db.Find(&employee)

	c.JSON(http.StatusOK, gin.H{"message": "Employee table created successfully"})


}


func GetEmploy(c *gin.Context){

	ID := c.Param("empid")

	db := database.DB

	var employee models.Employee

	db.Find(&employee, ID)

	c.JSON(http.StatusOK, gin.H{"message": "Employee table created successfully"})

}