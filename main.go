package main

import (
	"workshop/database"
	"workshop/router")

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run("localhost:8080")
}