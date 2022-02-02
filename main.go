package main

import (
	"be-payroll/api/router"
	"be-payroll/pkg/database"
	"log"
)

func main() {
	db, err := database.InitDB()

	if err != nil {
		log.Fatal("error connecting to database: ", err)
		return
	}

	app := router.Route(db)
	app.Run(":8081")

}
