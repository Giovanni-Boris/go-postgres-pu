package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-api-postgres/models"
    "go-api-postgres/database"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&user)
    c.JSON(http.StatusOK, user)
}
