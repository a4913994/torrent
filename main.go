package main

import (
	"downloader/config"
	"downloader/db"
	"downloader/models"
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
	models.Init()
	server.Init()
}
