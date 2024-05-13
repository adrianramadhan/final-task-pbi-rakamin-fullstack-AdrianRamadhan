package main

import (
	"github.com/adrianramadhan/final-task-pbi-rakamin-fullstack-AdrianRamadhan/database"
	"github.com/adrianramadhan/final-task-pbi-rakamin-fullstack-AdrianRamadhan/router"
)

func main() {
	database.ConnectDB()
	router := router.SetupRouter()

	router.Run(":8088")
}