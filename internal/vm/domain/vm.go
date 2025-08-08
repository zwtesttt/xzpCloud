package domain

type Vm struct {
	id           string
	name         string
	status       VmStatus
	userId       string
	ip           string
	config       *VmConfig
	createdAt    int64
	updatedAt    int64
	expirationAt int64
	deletedAt    int64
}

type VmConfig struct {
	cpu    int
	disk   string
	memory string
}

func NewVmConfig(cpu int, disk string, memory string) *VmConfig {
	return &VmConfig{cpu: cpu, disk: disk, memory: memory}
}

func (v *VmConfig) Cpu() int {
	return v.cpu
}

func (v *VmConfig) Disk() string {
	return v.disk
}

func (v *VmConfig) Memory() string {
	return v.memory
}

type VmStatus int

const (
	//VmStatusStop 停止
	VmStatusStop VmStatus = -1
	//VmStatusStart 运行中
	VmStatusStart VmStatus = 1
)

func NewVm(id string, name string, status VmStatus, userId string, ip string, cfg *VmConfig, createdAt int64, updatedAt int64, expirationAt int64, deletedAt int64) *Vm {
	return &Vm{id: id, name: name, status: status, userId: userId, ip: ip, config: cfg, createdAt: createdAt, updatedAt: updatedAt, expirationAt: expirationAt, deletedAt: deletedAt}
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

func (v *Vm) Ip() string {
	return v.ip
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

func (v *Vm) DeletedAt() int64 {
	return v.deletedAt
}

func (v *Vm) Config() *VmConfig {
	return v.config
}

func (v *Vm) SetStatus(status VmStatus) {
	v.status = status
}
