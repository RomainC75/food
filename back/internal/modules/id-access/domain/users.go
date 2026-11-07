package idaccess_dom

import (
	"context"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = "user not found"
	ErrCouldNotSaveUser = "could not save user"
)

type Users interface {
	Create(ctx context.Context, newUser *User) error
	FindById(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
