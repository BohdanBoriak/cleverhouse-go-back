package database

import "github.com/upper/db/v4"

type user struct {
	Id         uint64 `db:"id,omitempty"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Phone      string `db:"phone"`
	Email      string `db:"email"`
	Password   string `db:"password"`
}

func SaveBohdan(sess db.Session) error {
	u := user{
		FirstName:  "Bohdan",
		SecondName: "Boriak",
		Phone:      "+38000000000",
		Email:      "t@test.com",
		Password:   "12345678",
	}

	err := sess.
		Collection("users").
		InsertReturning(&u)

	return err
}
