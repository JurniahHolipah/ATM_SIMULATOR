package db

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

var DB *sql.DB

func Init() {
    var err error
   DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/atm_simulator")

    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Error pinging the database:", err)
    }

    log.Println("Database connected successfully!")
}
