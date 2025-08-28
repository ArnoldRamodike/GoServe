package main

import (
	"log"
	"os"

	"github.com/sikozonpc/goserve/internal/db"
	"github.com/sikozonpc/goserve/internal/env"
	"github.com/sikozonpc/goserve/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgress://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.Getint("DB_MAX_OPEN_CONS", 30),
			maxIdleConns: env.Getint("DB_MAX_OPEN_CONS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Printf("Database connection pool established!")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
