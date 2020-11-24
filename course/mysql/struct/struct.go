package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	//apos o "root: , COLOQUE A SENHA"
	db, err := sql.Open("mysql", "root:@/")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios(
			id integer auto_increment,
			nome varchar(30),
			PRIMARY KEY (id)
		)`)
}
