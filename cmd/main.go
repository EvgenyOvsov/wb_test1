package main

import (
	"draft/datastore"
	"draft/server"
	"os"
	"os/signal"
	"syscall"
)

// Дефолтная архитектура с обработкой прерываний
func main() {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	db := datastore.NewDatastore()
	srv := server.NewServer(db)
	go srv.Start()

	<-signalChan
	srv.Stop()
}
