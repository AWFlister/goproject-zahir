package main

import (
	"flag"
	seeder "goproject-zahir/db"
	"goproject-zahir/handler"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "reset" {
		os.Remove("db/zahir.db")
	}
	db, err := gorm.Open(sqlite.Open("db/zahir.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	h := handler.Handler{DB: db}
	h.Migrate()
	h.Register(r)

	if len(args) > 0 {
		switch args[0] {
		case "seed", "reset":
			seeder.Seeder{DB: db}.Seed()
		}
	}

	http.ListenAndServe(":8080", r)
}
