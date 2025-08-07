package domain

type Role struct {
	name        string
	description string
	policies    []*Policy
	roleType    RoleType
	createdAt   int64
	updatedAt   int64
}

type Policy struct {
	sid    string
	effect string
}

type RoleType int

const (
	RoleTypeUser  RoleType = iota // 用户
	RoleTypeAdmin                 // 管理员
)

func NewRole(name, description string, policies []*Policy, roleType RoleType, createdAt, updatedAt int64) *Role {
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

func (r *Role) RoleType() RoleType {
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
