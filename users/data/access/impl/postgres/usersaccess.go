package postgres

import (
	"bujem/users/model"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type UsersAccessPostgres struct{}

func (dao UsersAccessPostgres) Create(user *model.User) error {
	db, err := getConnection()

	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("insert into users (id, username, email) values ($1, $2, $3)", user.ID, user.Username, user.Email)

	if err != nil {
		return err
	}
	return nil
}

func (dao UsersAccessPostgres) Update(user *model.User) error {
	db, err := getConnection()

	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("update users set username=$1, email=$2 where id=$3", user.Username, user.Email, user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (dao UsersAccessPostgres) FindById(id int) (model.User, error) {
	db, err := getConnection()

	if err != nil {
		return model.User{}, err
	}

	defer db.Close()

	rows, err := db.Query("select id, username, email from users where id = $1", id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return model.User{}, fmt.Errorf("User record with ID %d not found", id)
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

func (dao UsersAccessPostgres) DeleteById(id int) error {
	db, err := getConnection()

	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("delete from users where id=$1", id)

	if err != nil {
		return err
	}
	return nil
}

func getConnection() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://postgres:postgres@localhost/bujem?sslmode=disable")
}
