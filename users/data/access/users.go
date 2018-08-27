package data

import (
	"bujem/users/model"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func FindUser(ID int) (model.User, error) {
	db, err := getConnection()

	if err != nil {
		return model.User{}, err
	}

	defer db.Close()

	rows, err := db.Query("select id, username, email from users where id = $1", ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return model.User{}, nil
	}

	var row model.User
	err = rows.Scan(&row.ID, &row.Username, &row.Email)
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}
	log.Println(row.ID, row.Username, row.Email)

	err = rows.Err()
	if err != nil {
		return model.User{}, err
	}

	return row, err
}

func getConnection() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://postgres:postgres@localhost/bujem?sslmode=disable")
}
