package main

import (
	"log"
	"net/http"

	// "os"
	"time"

	"github.com/mentalko/go_api_server/internal/repository"
	"github.com/mentalko/go_api_server/internal/service"
	"github.com/mentalko/go_api_server/internal/transport/rest"
	"github.com/mentalko/go_api_server/pkg/database"
)

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "postgres",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		Password: "pass",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	accountsRepo := repository.NewAccount(db)
	accountsService := service.NewAccount(accountsRepo)

	transactionsRepo := repository.NewTransaction(db)
	transactionsService := service.NewTransaction(transactionsRepo)
	handler := rest.NewHandler(accountsService, transactionsService)

	// init & run server
	srv := &http.Server{
		Addr:    ":80",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT ", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
