package main

import "github.com/apedorenkoS/gogger/cmd/config"

func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}
	InitGlobalLogger()
	NewServer().Start()
}
