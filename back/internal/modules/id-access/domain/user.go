package idaccess_dom

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	email    string
	password string
}

func NewUser(id uuid.UUID, email string, password string) *User {
	return &User{
		id:       id,
		email:    email,
		password: password,
	}
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) IsSameEmail(other string) bool {
	return u.email == other
}

func (u *User) Clone() *User {
	return &User{
		id:    u.id,
		email: u.email,
	}
}

// -----

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
	}{
		ID:       u.id,
		Email:    u.email,
		Password: u.password,
	})
}
