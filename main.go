package main

import (
    "github.com/gin-gonic/gin"
    "go-api-postgres/config"
    "go-api-postgres/database"
    "go-api-postgres/models"
    "go-api-postgres/handlers"
)

func main() {
    config.LoadEnv()
    config.LoadConfig()

    database.Connect()
    database.DB.AutoMigrate(&models.User{})

    r := gin.Default()
    r.GET("/users", handlers.GetUsers)
    r.POST("/users", handlers.CreateUser)

    r.Run(":"+config.PORT) // listen and serve on
}
