package main

import (
	"log"

	"github.com/nikkbh/social-app/internal/db"
	"github.com/nikkbh/social-app/internal/env"
	"github.com/nikkbh/social-app/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
