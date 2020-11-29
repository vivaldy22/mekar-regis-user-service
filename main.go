package main

import "github.com/vivaldy22/mekar-regis-user-service/config"

func main() {
	db, _ := config.InitDB()
	config.RunServer(db)
}
