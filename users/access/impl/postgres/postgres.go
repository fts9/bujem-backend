package postgres

import (
	"bujem/common/utility/sqlbuilder"
	"bujem/users/model"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // DB driver so is imported anonymously
)

// UsersAccessPostgres implements data access for users on a Postgres database
type UsersAccessPostgres struct{}

// Create adds a new user record to the database
func (dao UsersAccessPostgres) Create(user *model.User) error {
	db, err := getConnection()

	if err != nil {
		return err
	}

	var ID int64
	defer db.Close()

	// PostgreSQL doesn't support retrieving the last inserted ID from metadata so QueryRow and Scan act as a work around for this
	err = db.QueryRow("insert into users (username, email, password, created, modified) values ($1, $2, $3, now()::timestamp, now()::timestamp) RETURNING id", user.Username, user.Email, user.Password).Scan(&ID)
	if err != nil {
		return err
	}

	user.ID = ID
	return nil
}

// Update updates an existing user record on the database
func (dao UsersAccessPostgres) Update(user *model.User) error {
	db, err := getConnection()

	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("update users set username=$1, email=$2, password=$3, modified=now()::timestamp where id=$4", user.Username, user.Email, user.Password, user.ID)

	if err != nil {
		return err
	}
	return nil
}

// FindByID retrieves a user record by looking up the given ID
func (dao UsersAccessPostgres) FindByID(ID int64) (model.User, error) {
	selectQuery, err := sqlbuilder.NewSelectQueryWithLogging(true).Columns("*").Table("", "users").Where("id").Build()
	db, err := getConnection()

	if err != nil {
		return model.User{}, err
	}

	defer db.Close()

	rows, err := db.Query(selectQuery, ID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return model.User{}, fmt.Errorf("User record with ID %d not found", ID)
	}

	var row model.User
	err = rows.Scan(&row.ID, &row.Username, &row.Email, &row.Password, &row.Created, &row.Modified)
	if err != nil {
		log.Println(err)
		return model.User{}, err
	}
	log.Println(row.ID, row.Username, row.Email)

	err = rows.Err()
	if err != nil {
		return model.User{}, err
	}

	return row, nil
}

// DeleteByID deletes a user record by looking up the given ID
func (dao UsersAccessPostgres) DeleteByID(ID int64) error {
	db, err := getConnection()

	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("delete from users where id=$1", ID)

	if err != nil {
		return err
	}
	return nil
}

func getConnection() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://postgres:postgres@localhost/bujem?sslmode=disable")
}
