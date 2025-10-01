package main

import (
	"github.com/nepile/event-scanner/infrastructure/config"
	"github.com/nepile/event-scanner/infrastructure/db"
)

func main() {
	config.LoadConfig()
	db.ConnectDB()
}
