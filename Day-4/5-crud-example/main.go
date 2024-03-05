package main

import (
	"fmt"

	"github.com/marlon-gaules/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("meubanco.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&users.User{})

	id := users.Create(db, "Joaquim")
	lista := users.List(db)
	for _, user := range lista {
		fmt.Println(user.Name)
	}
	users.Update(db, id, "João")
	users.Create(db, "Joaquim 1")
	users.Create(db, "Joaquim 2")
	users.Create(db, "Joaquim 3 ")
	lista = users.List(db)
	for _, user := range lista {
		fmt.Println(user.Name)
	}
	users.Delete(db, id)
	lista = users.List(db)
	for _, user := range lista {
		fmt.Println(user.Name)
	}
}
