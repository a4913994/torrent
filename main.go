package main

import (
	"downloader/config"
	"downloader/db"
	"downloader/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*env)
	db.Init()
	server.Init()
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Welcome Gin Server")
// 	})

// 	srv := &http.Server{
// 		Addr:    ":8080",
// 		Handler: r,
// 	}

// 	// Initialzing the server in a gorutine so that
// 	// it won't block the graceful shutdown handling below
// 	go func() {
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("listen: %s\n", err)
// 		}
// 	}()

// 	// Wait for interrupt signal to gracefully shutdown the server with
// 	// a timeout of 5 seconds
// 	quit := make(chan os.Signal)
// 	// kill (no param) default send syscall.SIGTERM
// 	// kill -2 is syscall.SIGIN
// 	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// 	log.Println("Shutting down server ....")

// 	// The context is used to inform the server it has 5 seconds to finish
// 	// the request it is current handing
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := srv.Shutdown(ctx); err != nil {
// 		log.Fatal("Server forced to shutdown:", err)
// 	}

// 	log.Println("Server exting")
// }
