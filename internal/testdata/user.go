package testdata

import (
	"time"
)

type EntityID struct {
	ID string `db:"fqn"`
}

type EntityDeletedAt struct {
	DeletedAt time.Time `db:"deleted_at"`
}

type Entity struct {
	ID        string    `db:"uuid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	EntityID
	EntityDeletedAt
}

type User struct {
	Entity

	ID        int64  `db:"id"`
	FirstName string `db:"name"`
	LastName  string `db:"surname"`
	Email     string `db:"email"`
	Password  string `db:"password"`

	FullName string
}
