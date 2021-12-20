package main

import (
	"fmt"

	"github.com/KeshikaGupta20/Postgresql_GO/database"
	"github.com/KeshikaGupta20/Postgresql_GO/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	r := gin.Default()

	godotenv.Load()

	database.ConnectDB()

	fmt.Println("Server started at port 4000")

	router.RegisterRoutes(r)

	r.Run(":3000")
}
