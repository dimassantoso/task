package main

import (
	"fmt"
	configEnv "github.com/joho/godotenv"
	"microservice-assignment/app"
	"microservice-assignment/config"
	"microservice-assignment/shared"
	"os"
)

func main() {
	err := configEnv.Load(".env")
	if err != nil {
		os.Exit(2)
	}

	//initialize database,  only once for entire services
	mySQLReadDB := config.ReadMySQLDB()
	mySQLWriteDB := config.WriteMySQLDB()

	// set schema
	errJson := shared.Load(shared.DirectorySchema)
	if errJson != nil {
		fmt.Println("Error json schema load")
		fmt.Println(errJson)
		os.Exit(2)
	}

	run := app.MakeHandler(mySQLReadDB, mySQLWriteDB)
	run.Start()

}
