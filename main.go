package main

import "github.com/Panitnun-6243/duckduck-server/db"

func main() {
	db.InitializeDB()
	defer db.DisconnectDB()
}
