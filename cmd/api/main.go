package main

import (
	"log"
	"social-api/internal/env"
	store "social-api/internal/storage"
)

func main() {

    cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgresql://postgres:laudapassword@db.notvibecoderbitches.supabase.co:5432/postgres"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_TIME", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
