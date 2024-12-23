package domain

import "github.com/zwtesttt/xzpCloud/pkg/role"

// User
type User struct {
	id        string
	name      string
	email     string
	password  string
	avatar    string
	roleId    role.Type
	createdAt int64
	updatedAt int64
	deletedAt int64
}

func NewUser(id, name, email, password, avatar string, roleId role.Type, createdAt, updatedAt, deletedAt int64) *User {
	return &User{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		avatar:    avatar,
		roleId:    roleId,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Avatar() string {
	return u.avatar
}

func (u *User) RoleId() role.Type {
	return u.roleId
}

func (u *User) CreatedAt() int64 {
	return u.createdAt
}

func (u *User) UpdatedAt() int64 {
	return u.updatedAt
}

func (u *User) DeletedAt() int64 {
	return u.deletedAt
}
