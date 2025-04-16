package repository

import (
	"context"
	"github.com/uptrace/bun"
	"strings"
)

type User struct {
	Iin        string `bun:"iin"`
	Firstname  string `bun:"firstname"`
	Lastname   string `bun:"lastname"`
	Middlename string `bun:"middlename"`
	Phone      string `bun:"phone"`
}

type userRepo struct {
	db bun.IDB
}

type UserSDK interface {
	Create(ctx context.Context, request *User) error
	ListUserByName(ctx context.Context, name string) ([]*User, error)
	GetUserByIin(ctx context.Context, iin string) (*User, error)
}

func NewUserRepo(db bun.IDB) UserSDK {
	return &userRepo{db: db}
}

func (u *userRepo) Create(ctx context.Context, info *User) error {
	info.Firstname = strings.ToLower(info.Firstname)
	info.Lastname = strings.ToLower(info.Lastname)
	info.Middlename = strings.ToLower(info.Middlename)
	_, err := u.db.NewInsert().Model(info).Exec(ctx)
	return err
}

func (u *userRepo) ListUserByName(ctx context.Context, name string) ([]*User, error) {
	name = strings.ToLower(name)

	query := u.db.NewSelect().Model((*User)(nil))

	query = query.
		WhereOr("LOWER(firstname) LIKE ?", name+"%").
		WhereOr("LOWER(lastname) LIKE ?", name+"%").
		WhereOr("LOWER(middlename) LIKE ?", name+"%")

	var users []*User
	err := query.Scan(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (u *userRepo) GetUserByIin(ctx context.Context, iin string) (*User, error) {
	var user User
	err := u.db.NewSelect().
		Model(&user).
		Where("iin = ?", iin).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
