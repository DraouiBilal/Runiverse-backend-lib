package main

import (
	database "github.com/DraouiBilal/Runiverse-backend-lib/db"
	"log"
)

type Test struct {
	Id    string
	Age   int
	Name  string
	Email string
}

func main() {

	db, err := database.ConnectDB("0.0.0.0", "postgres", "postgres", "mydb", "disable", 5432)

	if err == nil {
		log.Println("DB connected")
	}else {
		log.Fatal("Can't connect to the DB")	
	}

	_, db_err := database.CreateTable(db, "test", map[string][]string{
		"id":    {"VARCHAR(50)", "PRIMARY KEY"},
		"name":  {"VARCHAR(200)"},
		"age":   {"INT"},
		"email": {"VARCHAR(50)", "UNIQUE"},
	})

	if db_err != nil {
		log.Println(db_err)
	}

	res, query_err := database.QueryAll[Test](db, "test")
	
	if query_err != nil {
		log.Println(query_err)
	}

	for _,val := range res {
		log.Println(val)
	}

}
