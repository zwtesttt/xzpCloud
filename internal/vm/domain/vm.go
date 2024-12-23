package domain

type Vm struct {
	id           string
	name         string
	status       VmStatus
	userId       string
	createdAt    int64
	updatedAt    int64
	expirationAt int64
}

type VmStatus int

const (
	//VmStatusStop 停用
	VmStatusStop VmStatus = -1
	//VmStatusStart 正常
	VmStatusStart VmStatus = 1
)

func NewVm(id string, name string, status VmStatus, userId string, createdAt int64, updatedAt int64, expirationAt int64) *Vm {
	return &Vm{id: id, name: name, status: status, userId: userId, createdAt: createdAt, updatedAt: updatedAt, expirationAt: expirationAt}
}

func (v *Vm) Id() string {
	return v.id
}

func (v *Vm) Name() string {
	return v.name
}

func (v *Vm) Status() VmStatus {
	return v.status
}

func (v *Vm) UserId() string {
	return v.userId
}

func (v *Vm) CreatedAt() int64 {
	return v.createdAt
}

func (v *Vm) UpdatedAt() int64 {
	return v.updatedAt
}

func (v *Vm) ExpirationAt() int64 {
	return v.expirationAt
}
