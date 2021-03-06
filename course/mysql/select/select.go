package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	// Depois do "root:" tem que colocar a senha
	db, err := sql.Open("mysql", "root:@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}