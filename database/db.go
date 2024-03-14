package database

import (
    "fmt"
    "log"
    "os"
    "database/sql"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

var (
    DB  *sql.DB
)

func StartDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    host := os.Getenv("HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    config := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    DB, err = sql.Open("postgres", config)
    if err != nil {
        panic(err)
    }

    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected to database")
}
