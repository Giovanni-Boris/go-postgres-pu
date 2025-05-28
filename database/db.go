package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "go-api-postgres/config"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=" + config.DBHost +
        " user=" + config.DBUser +
        " password=" + config.DBPassword +
        " dbname=" + config.DBName +
        " port=" + config.DBPort +
        " sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    DB = db
}
