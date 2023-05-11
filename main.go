package main

import (
	"fmt"

	"airland-server/src/config"
	"airland-server/src/cross_cutting/ioc"
)

func main() {
	config.ReadConfig(".")
	fmt.Println("environment variables loaded successfully...")

	server, err := ioc.Bootstrap()

	if err != nil {
		panic(err)
	}
	fmt.Println("dependencies injected successfully")
	server.Run()
}
