package repository

import "github.com/VitorC-Sa/GolangCleanArchitecture/app/entity"

type UserRepository interface {
	Insert(user *entity.User) error
	Get(searchBy string, value interface{}) (*entity.User, error)
	Update(user *entity.User) error
	Delete(deleteBy string, value interface{}) error
}
