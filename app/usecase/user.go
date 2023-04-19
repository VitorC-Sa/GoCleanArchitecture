package usecase

import "github.com/VitorC-Sa/GolangCleanArchitecture/app/entity"

type UserRepository interface {
	Insert(user *entity.User) error
	Get(searchBy string, value interface{}) (*entity.User, error)
	Update(user *entity.User) error
	Delete(deleteBy string, value interface{}) error
}

type UserUseCase struct {
	User UserRepository
}

func NewUserUseCase(ur UserRepository) UserUseCase {
	return UserUseCase{User: ur}
}

func (uu *UserUseCase) AddUser(user *entity.User) error {
	return uu.User.Insert(user)
}

func (uu *UserUseCase) GetUser(searchBy string, value interface{}) (*entity.User, error) {
	return uu.User.Get(searchBy, value)
}

func (uu UserUseCase) UpdateUser(user *entity.User) error {
	return uu.User.Update(user)
}

func (uu UserUseCase) DeleteUser(deleteBy string, value interface{}) error {
	return uu.User.Delete(deleteBy, value)
}
