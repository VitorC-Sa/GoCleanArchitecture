package repository

import (
	"context"
	"database/sql"

	"github.com/VitorC-Sa/GolangCleanArchitecture/app/entity"
	"github.com/uptrace/bun"
)

type userRepositoryBun struct {
	db *bun.DB
}

func NewUserRepositoryBun(db *bun.DB) UserRepository {
	return &userRepositoryBun{db: db}
}

func (ur *userRepositoryBun) Insert(user *entity.User) error {
	_, err := ur.db.
		NewInsert().
		Model(user).
		Exec(context.TODO())
	return err
}

func (ur *userRepositoryBun) Get(searchBy string, value interface{}) (*entity.User, error) {
	user := new(entity.User)
	err := ur.db.
		NewSelect().
		Model(user).
		Where("? = ?", bun.Ident(searchBy), value).
		Scan(context.TODO())

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

func (ur *userRepositoryBun) Update(user *entity.User) error {
	_, err := ur.db.
		NewUpdate().
		Model(user).
		WherePK().
		Exec(context.TODO())
	return err
}

func (ur *userRepositoryBun) Delete(deleteBy string, value interface{}) error {
	_, err := ur.db.
		NewDelete().
		Model(&entity.User{}).
		Where("? = ?", bun.Ident(deleteBy), value).
		Exec(context.TODO())

	return err
}
