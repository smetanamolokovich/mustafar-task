package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/smetanamolokovich/mustafar_task/pkg/kvstore"
)

type application struct {
	store  *kvstore.KvStore
	logger *log.Logger
}

func main() {
	port := os.Getenv("PORT")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := connectDB(os.Getenv("DB_DSN"))
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Println("DB connection established")

	store := kvstore.New(db, "kvstore")

	app := &application{
		store:  store,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Server started on %s", srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func connectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
