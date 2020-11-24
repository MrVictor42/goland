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

	tx, _ := db.Begin()
	/*
		Como na linha acima se começa a transação, todo o código abaixo até o commit() não irá funcionar, pois vai da erro no exec do thiago, por que já existe o id = 1 salvo no banco de dados, então tem o problema de chave duplicada.	
	*/
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Thiago") // chave duplicada

	/*
		A linha abaixo é justamente para tratar o erro da chave duplicada acima, pois como irá da erro, toda a transação é cancelada e o Carlos e a Bia não são adicionados.
	*/
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
