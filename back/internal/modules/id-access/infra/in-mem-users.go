package idaccess_infra

import (
	"context"
	"errors"
	idaccess_dom "food/internal/modules/id-access/domain"

	"github.com/google/uuid"
)

type InMemoryUsers struct {
	savedUsers                []*idaccess_dom.User
	isSavingUserErrorExpected bool
	expectedUser              *idaccess_dom.User
	deletedUser               *idaccess_dom.User
}

func NewInMemoryUsers() *InMemoryUsers {
	return &InMemoryUsers{
		savedUsers: make([]*idaccess_dom.User, 0, 10),
	}
}

func (imu *InMemoryUsers) Create(ctx context.Context, newUser *idaccess_dom.User) error {
	if imu.isSavingUserErrorExpected {
		return errors.New(idaccess_dom.ErrCouldNotSaveUser)
	}
	imu.savedUsers = append(imu.savedUsers, newUser)
	return nil
}

func (imu *InMemoryUsers) FindById(ctx context.Context, id uuid.UUID) (*idaccess_dom.User, error) {
	for _, u := range imu.savedUsers {
		if u.Id() == id {
			return u.Clone(), nil
		}
	}
	return nil, errors.New(idaccess_dom.ErrUserNotFound)
}

func (imu *InMemoryUsers) FindByEmail(ctx context.Context, email string) (*idaccess_dom.User, error) {
	for _, u := range imu.savedUsers {
		if u.IsSameEmail(email) {
			return u.Clone(), nil
		}
	}
	return nil, errors.New(idaccess_dom.ErrUserNotFound)
}

// ========= TESTS =========

func (imu *InMemoryUsers) SetupIsSavingUserError(isSavingUserErrorExpected bool) {
	imu.isSavingUserErrorExpected = isSavingUserErrorExpected
}

func (imu *InMemoryUsers) SavedUsers() []*idaccess_dom.User {
	return imu.savedUsers
}
