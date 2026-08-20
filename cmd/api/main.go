package main

import (
	"log"
	"social-api/internal/env"
	store "social-api/internal/storage"
)

func main() {

    cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
