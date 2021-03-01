package main

import (
	"github.com/disebud/api-users-test/config"
	"github.com/disebud/api-users-test/main/master"
)

func main() {
	db, _ := config.Connection()    //ok
	router := config.CreateRouter() //ok
	master.Init(router, db)
	config.RunServer(router)
}
