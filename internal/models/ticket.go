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
	OrderNum     string
	Status       string // running/passed/rejected
	uid          string
	Step         string
	Operator     []string
	OperatedUser []string // 在Disposal.SignType为jointly_sign/serial_sign时使用
	Memo         string
}

type Disposal struct {
	SignType      string  // jointly_sign/serial_sign/anyone_sign
	JointSignRate float32 // 仅jointly_sign时使用
}

// 发起工单时 可以直接使用模版 或者自定义模版 自定义模版需要
type TicketTemplate struct {
	uid       string
	EndStep   []string        // 结束节点
	StartStep string          // 开始节点
	Config    []*TicketConfig // 配置
	Builtin   bool            // 是否内置
}

type TicketConfig struct {
	Step     string      // 步骤名
	State    string      // 步骤所属状态
	Operator []string    // 预设操作人
	Next     []*NextStep // 下一节点
	Disposal Disposal    // 处置方式
}

func (t *TicketConfig) GetNext() []*NextStep {
	return utils.TernaryOperator(t == nil, nil, t.Next)
}

type NextStep struct {
	Step      string // 步骤名
	Operation string // 操作名
}

func (n *NextStep) GetStep() string {
	return utils.TernaryOperator(n == nil, "", n.Step)
}

func (n *NextStep) GetOperation() string {
	return utils.TernaryOperator(n == nil, "", n.Operation)
}
