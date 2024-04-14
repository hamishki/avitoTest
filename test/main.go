package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	SQLOpen()
}

func SQLOpen() {
	IsItAliveQuestionMark(context.Background())

}

func IsItAliveQuestionMark(ctx context.Context) {
	// db, err := sql.Open("pgx", "habrpguser://pgx_md5:pgpwd4habr@localhost:5432/habrdb")
	db, err := sql.Open("pgx", "user=gopher password=pass host=localhost port=5432 database=test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	Query(ctx, db)
}

func Query(ctx context.Context, db *sql.DB) {
	rows, err := db.QueryContext(ctx, "SELECT RoleID, RoleName FROM Role")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		log.Println(id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

//docker run --name habr-pg-13.3 -p 5432:5432 -e POSTGRES_USER=habrpguser -e POSTGRES_PASSWORD=pgpwd4habr
//-e POSTGRES_DB=habrdb -d postgres:13.3

func Exec(ctx context.Context, db *sql.DB) {
	res, err := db.ExecContext(
		ctx,
		"UPDATE Role SET RoleName = xd WHERE RoleID = $1",
	)
	if err != nil {
		log.Fatal(err)
	}

	lastID, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()

	log.Println(lastID, rowsAffected)
}
