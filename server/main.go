package main

import (
	server "github.com/LevonAsatryan/feature-flags/api"
)

func main() {
	var port int64 = 3001
	ser := server.NewServer(port)
	err := ser.Start()

	if err != nil {
		panic(err)
	}
}
