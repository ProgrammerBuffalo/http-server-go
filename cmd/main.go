package main

import (
	"context"
	"http-app/handlers"
	"http-app/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/lib/pq"
)

func startServer(wg *sync.WaitGroup, server *http.Server) {
	log.Println("goroutine1: Start server...")
	defer wg.Done()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func shutdownServer(wg *sync.WaitGroup, server *http.Server) {
	log.Println("goroutine2: Check interrput OS signal...")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	<-ctx.Done()

	defer wg.Done()
	defer stop()
	defer log.Println("goroutine2: Disconnect server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/accounts", handlers.GetAccounts)

	server := &http.Server{
		Addr:    ":8080",
		Handler: serverMux,
	}

	wg := new(sync.WaitGroup)

	internal.DBConnect()

	wg.Add(2)

	go startServer(wg, server)
	go shutdownServer(wg, server)

	log.Println("Main thread wait...")
	wg.Wait()
	log.Println("Main thread start...")

	internal.DBDisconnect()

	log.Println("Server successfully finished...")

}
