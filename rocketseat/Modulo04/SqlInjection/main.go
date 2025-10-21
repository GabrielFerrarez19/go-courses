package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}

	createTableSql := `
	CREATE TABLE foo(
	id integer not null primary key,
	name text
	)
	`

	res, err := db.Exec(createTableSql)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	insertSql := `
	INSERT INTO foo (id, name) values (1,"pedro")
	`

	res, err = db.Exec(insertSql)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RowsAffected())

	type user struct {
		ID   int64
		Name string
	}

	querySql := `
	SELECT * FROM foo WHERE ID = ?
	`

	var u user
	if err := db.QueryRow(querySql, 1).Scan(&u.ID, &u.Name); err != nil {
		panic(err)
	}

	fmt.Println(u)

	for id := 2; id < 10; id++ {
		_, err := db.Exec("INSERT INTO foo (id) values (?)", id)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("------antes-------")

	var count int64

	if err := db.QueryRow("SELECT count(*) from foo;").Scan(&count); err != nil {
		panic(err)
	}

	fmt.Println(count)

	input := "1 OR 1 = 1"
	// deleteSql := fmt.Sprintf(`
	// DELETE FROM foo WHERE id = %s
	// `, input)

	deleteSql := `
	DELETE FROM foo WHERE id = ?
	`

	if _, err := db.Exec(deleteSql, input); err != nil {
		panic(err)
	}

	fmt.Println("------depois-------")

	if err := db.QueryRow("SELECT count(*) from foo;").Scan(&count); err != nil {
		panic(err)
	}

	fmt.Println(count)
}
