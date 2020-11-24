package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Depois do "root:" tem que colocar a senha
	db, err := sql.Open("mysql", "root:@/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uóxition Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	// delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}