package idaccess_infra

import (
	"context"
	idaccess_dom "food/internal/modules/id-access/domain"
	"log"
	"sync"

	"github.com/google/uuid"
)

type InMemoryUsers struct {
	m                         sync.RWMutex
	savedUsers                map[uuid.UUID]*idaccess_dom.User
	isSavingUserErrorExpected bool
	expectedUser              *idaccess_dom.User
	deletedUser               *idaccess_dom.User
}

func NewInMemoryUsers() *InMemoryUsers {
	return &InMemoryUsers{
		m:          sync.RWMutex{},
		savedUsers: make(map[uuid.UUID]*idaccess_dom.User, 10),
	}
}

func (imu *InMemoryUsers) Create(ctx context.Context, newUser *idaccess_dom.User) error {
	if imu.isSavingUserErrorExpected {
		return idaccess_dom.ErrCouldNotSaveUser
	}
	imu.m.Lock()
	imu.savedUsers[newUser.Id()] = newUser.Clone()
	imu.m.Unlock()

	imu.DisplayAll()
	return nil
}

func (imu *InMemoryUsers) FindById(ctx context.Context, id uuid.UUID) (*idaccess_dom.User, error) {
	imu.m.RLock()
	found, ok := imu.savedUsers[id]
	imu.m.RUnlock()
	if !ok {
		return nil, idaccess_dom.ErrUserNotFound
	}
	return found.Clone(), nil
}

func (imu *InMemoryUsers) FindByEmail(ctx context.Context, email string) (*idaccess_dom.User, error) {
	imu.m.RLock()
	defer imu.m.RUnlock()
	for _, u := range imu.savedUsers {
		if u.Email() == email {
			return u.Clone(), nil
		}
	}
	return nil, idaccess_dom.ErrUserNotFound
}

func (imu *InMemoryUsers) Init(user *idaccess_dom.User) {
	imu.savedUsers[user.Id()] = user.Clone()
}

// ========= TESTS =========

func (imu *InMemoryUsers) SetupIsSavingUserError(isSavingUserErrorExpected bool) {
	imu.isSavingUserErrorExpected = isSavingUserErrorExpected
}

func (imu *InMemoryUsers) SavedUsers() map[uuid.UUID]*idaccess_dom.User {
	return imu.savedUsers
}

func (imu *InMemoryUsers) DisplayAll() {
	log.Println("------- DISPLAY USERS --------")
	for _, user := range imu.savedUsers {
		u, _ := user.MarshalJSON()
		log.Println(string(u))
	}
	log.Println("------- END ---------")
}
