package models

import "github.com/FatWang1/fatwang-go-utils/utils"

const (
	JointlySign = "jointly_sign"
	SerialSign  = "serial_sign"
	AnyoneSign  = "anyone_sign"

	Running  = "running"
	Passed   = "passed"
	Rejected = "rejected"

	Reject = "reject"
)

type Ticket struct {
	OrderNum     string   `json:"order_num" gorm:"column:order_num;type:varchar(255);not null"`
	Status       string   `json:"status" gorm:"column:status;type:varchar(50);not null;default:'running'"` // running/passed/rejected
	Uid          string   `json:"uid" gorm:"column:uid;type:varchar(255);not null;primaryKey"`
	Step         string   `json:"step" gorm:"column:step;type:varchar(255);not null"`
	Operator     []string `json:"operator" gorm:"column:operator;type:json"`
	OperatedUser []string `json:"operated_user" gorm:"column:operated_user;type:json"` // 在Disposal.SignType为jointly_sign/serial_sign时使用
	Memo         string   `json:"memo" gorm:"column:memo;type:text"`
}

type Disposal struct {
	SignType      string  `json:"sign_type" gorm:"column:sign_type;type:varchar(50);not null"`        // jointly_sign/serial_sign/anyone_sign
	JointSignRate float32 `json:"joint_sign_rate" gorm:"column:joint_sign_rate;type:float;default:0"` // 仅jointly_sign时使用
}

// 发起工单时 可以直接使用模版 或者自定义模版 自定义模版需要
type TicketTemplate struct {
	Uid       string        `json:"uid" gorm:"column:uid;type:varchar(255);not null;primaryKey"`
	EndStep   []string      `json:"end_step" gorm:"column:end_step;type:json"`                      // 结束节点
	StartStep string        `json:"start_step" gorm:"column:start_step;type:varchar(255);not null"` // 开始节点
	Config    []*StepConfig `json:"config" gorm:"column:config;type:json"`                          // 配置
	Builtin   bool          `json:"builtin" gorm:"column:builtin;type:boolean;default:false"`       // 是否内置
}

type StepConfig struct {
	Step     string      `json:"step" gorm:"column:step;type:varchar(255);not null"`   // 步骤名
	State    string      `json:"state" gorm:"column:state;type:varchar(255);not null"` // 步骤所属状态
	Operator []string    `json:"operator" gorm:"column:operator;type:json"`            // 预设操作人
	Next     []*NextStep `json:"next" gorm:"column:next;type:json"`                    // 下一节点
	Disposal Disposal    `json:"disposal" gorm:"column:disposal;type:json"`            // 处置方式
}

func (t *StepConfig) GetNext() []*NextStep {
	return utils.TernaryOperator(t == nil, nil, t.Next)
}

type NextStep struct {
	Step      string `json:"step" gorm:"column:step;type:varchar(255);not null"`           // 步骤名
	Operation string `json:"operation" gorm:"column:operation;type:varchar(255);not null"` // 操作名
}

func (n *NextStep) GetStep() string {
	return utils.TernaryOperator(n == nil, "", n.Step)
}

func (n *NextStep) GetOperation() string {
	return utils.TernaryOperator(n == nil, "", n.Operation)
}
