package main

import (
	"context"
	"http-app/handlers"
	"http-app/storage"
	"http-app/storage/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func startServer(wg *sync.WaitGroup, server *http.Server) {
	log.Println("Start listen server...")
	defer wg.Done()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
	log.Println("End listen server...")
}

func shutdownServer(wg *sync.WaitGroup, server *http.Server) {
	log.Println("Start checking os signals...")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	<-ctx.Done()
	log.Println("End checking os signals...")

	defer wg.Done()
	defer log.Println("Shutdown server...")
	defer stop()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
}

func main() {
	serverMux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}

	wg := new(sync.WaitGroup)

	// Initialize db connection
	dbConnection := storage.DBInit()

	// Create repository
	accRepository := repository.NewAccountRepository(dbConnection)

	// Create handler and inject repository
	accHandler := handlers.NewAccountHandler(accRepository)

	// Set url for specific handler
	serverMux.HandleFunc("/accounts", accHandler.GetAccounts)

	wg.Add(2)

	go startServer(wg, server)
	go shutdownServer(wg, server)

	wg.Wait()

	storage.DBClose(dbConnection)

	log.Println("Server successfully finished...")
}
