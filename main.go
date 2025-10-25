package main

import (
	"log"
	"github.com/DraouiBilal/Runiverse-backend-lib/service"
)

type Test struct {
	Id    string
	Age   int
	Name  string
	Email string
}

func main() {

	id := service.GenerateID()
	log.Println("Generated ID:", id)
}
