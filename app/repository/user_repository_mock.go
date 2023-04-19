package repository

import (
	"github.com/VitorC-Sa/GolangCleanArchitecture/app/entity"
	"github.com/google/uuid"
)

type userRepositoryMock struct{}

var mockUser *entity.User

func NewUserRepositoryMock() UserRepository {
	return &userRepositoryMock{}
}

func (userRepositoryMock) Insert(user *entity.User) error {
	mockUser = user
	mockUser.ID = uuid.New().String()

	return nil
}

func (userRepositoryMock) Get(searchBy string, value interface{}) (*entity.User, error) {
	return mockUser, nil
}

func (userRepositoryMock) Update(user *entity.User) error {
	mockUser = user
	return nil
}

func (userRepositoryMock) Delete(deleteBy string, value interface{}) error {
	mockUser = nil
	return nil
}
