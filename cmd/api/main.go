package main

import (
	"log"
	"os"

	"github.com/sikozonpc/goserve/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}

	app := &application{
		config: cfg,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
