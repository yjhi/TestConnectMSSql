package main

import (
	"fmt"
)

func main() {

	cfg, err := LoadConfig("config.json")

	if err != nil {
		fmt.Println(err.Error())
		select {}
	}

	fmt.Println("Config:")
	fmt.Println(string(FileData))

	go StartWork(cfg)

	select {}

}
