package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/lib/pq"
)

// User is an abstraction of a single Users DB record
type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
	Modified NullTime  `json:"modified"`
}

type NullTime pq.NullTime // TODO How am I going to make this DBMS agnostic?

func (nt *NullTime) Scan(value interface{}) error {
	var timeStamp pq.NullTime
	if err := timeStamp.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*nt = NullTime{timeStamp.Time, false}
	} else {
		*nt = NullTime{timeStamp.Time, true}
	}

	return nil
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))), nil
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {
	s := string(b)

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}
