package server

import (
	"context"
	"downloader/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	srv := &http.Server{
		Addr:    c.GetString("server.port"),
		Handler: r,
	}

	// Initialzing the server in a gorutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGIN
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ....")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is current handing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
