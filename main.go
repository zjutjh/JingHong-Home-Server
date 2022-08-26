package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	. "zjutjh/Join-Us/config"
	_ "zjutjh/Join-Us/db"
	. "zjutjh/Join-Us/router"
	JHLog "zjutjh/Join-Us/utility/log"
)

func main() {
	var server *http.Server
	var port string = ":" + Config.Server.Port

	log.SetOutput(JHLog.LogMultiWriter)

	log.Println("Running Server at", port)
	server = &http.Server{
		Addr:    port,
		Handler: Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server Exiting")
}
