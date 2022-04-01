package user

import (
	"time"

	"github.com/google/uuid"
)

type UserServicer interface {
	// Register(newUser NewUser) (User, error)
	New() (NewAnonymousUser, error)
}

type userService struct {
	accessor UserAccessor
}

type UserAccessor interface {
	// CreateUser(newUser NewUser) (User, error)
	CreateAnonymousUser(user NewAnonymousUser) error
	// UpdateUser(user User) error
	// RemoveUser(user User) error
}

func NewUserService(userAccessor UserAccessor) UserServicer {
	return &userService{
		accessor: userAccessor,
	}
}

func (us *userService) New() (NewAnonymousUser, error) {
	user := NewAnonymousUser{
		UserId:       uuid.New().String(),
		CreatedOn:    time.Now(),
		LastActivity: time.Now(),
	}

	err := us.accessor.CreateAnonymousUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
