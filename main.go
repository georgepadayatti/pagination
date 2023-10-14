package main

import (
	"fmt"
	"log"

	"github.com/georgepadayatti/pagination/cmd"
	"github.com/georgepadayatti/pagination/db"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	log.Println("Database session opened")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
