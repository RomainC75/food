package idaccess_dom

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	id            uuid.UUID
	email         string
	password      string
	workspacesIDs []uuid.UUID
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
	if u == nil {
		return nil
	}
	return &User{
		id:       u.id,
		email:    u.email,
		password: u.password,
	}
}

// -----

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID    uuid.UUID `json:"id"`
		Email string    `json:"email"`
	}{
		ID:    u.id,
		Email: u.email,
	})
}
