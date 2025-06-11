// cmd/main.go
package main

import (
	"log"
	"turcompany/internal/app"
	"turcompany/db"
)

func main() {
	app.Run()
    db, err := db.ConnectDB()
    if err != nil {
        log.Fatalf("DB connection error: %v", err)
    }
    defer db.Close()

    // передаем db в router, сервис и т.д.
}
