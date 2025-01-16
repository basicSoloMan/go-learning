package main

import (
	"go-learning/internal/env"
	store2 "go-learning/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", "localhost:8080"),
	}

	store := store2.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
