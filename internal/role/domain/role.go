package domain

import "github.com/zwtesttt/xzpCloud/pkg/role"

type Role struct {
	name        string
	description string
	policies    []*Policy
	roleType    role.Type
	createdAt   int64
	updatedAt   int64
}

type Policy struct {
	sid    string
	effect string
}

func NewRole(name, description string, policies []*Policy, roleType role.Type, createdAt, updatedAt int64) *Role {
	return &Role{
		name:        name,
		description: description,
		policies:    policies,
		roleType:    roleType,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func NewPolicy(sid, effect string) *Policy {
	return &Policy{
		sid:    sid,
		effect: effect,
	}
}

func (r *Role) Name() string {
	return r.name
}

func (r *Role) Description() string {
	return r.description
}

func (r *Role) Policies() []*Policy {
	return r.policies
}

func (r *Role) RoleType() role.Type {
	return r.roleType
}

func (r *Role) CreatedAt() int64 {
	return r.createdAt
}

func (r *Role) UpdatedAt() int64 {
	return r.updatedAt
}

func (p *Policy) Sid() string {
	return p.sid
}

func (p *Policy) Effect() string {
	return p.effect
}
