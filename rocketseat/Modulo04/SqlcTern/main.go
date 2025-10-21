package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	urlExample := "postgres://root:12345@localhost:5432/tests"
	db, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	queries := New(db)
	ctx := context.Background()

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(authors)

	author, err := queries.CreateAuthor(ctx, CreateAuthorParams{
		Name: "gabriel ferrarez",
		Bio:  pgtype.Text{String: "Programador fulstack", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(author)
}
