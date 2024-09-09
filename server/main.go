package main

import (
	"github.com/kjunmin/online-judge/backend/api"
)

func main() {
	server := api.InitApiServer(":8080")
	if err := server.Run(); err != nil {
		panic(err)
	}
}
