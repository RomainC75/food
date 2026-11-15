package idaccess_dom

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrCouldNotSaveUser = errors.New("could not save user")
)

type Users interface {
	Create(ctx context.Context, newUser *User) error
	FindById(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
