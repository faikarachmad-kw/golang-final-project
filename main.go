package main

import (
	"final-project/database"
	"final-project/router"
)

func main() {
	const PORT =":8080"
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
