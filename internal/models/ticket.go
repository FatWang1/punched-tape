package models

type Ticket struct {
	OrderNum string
	Status   string // running/pass/rejected
	uid      string
	Step     string
	Operator []string
	Memo     string
}

type Disposal struct {
	SignType      string  // jointly_sign/serial_sign/anyone_sign
	JointSignRate float32 // 仅jointly_sign时使用
}

type TicketTemplate struct {
	uid       string
	EndStep   []string
	StartStep string
	Config    []*TicketConfig
}

type TicketConfig struct {
	Step     string
	State    string
	Operator []string
	Next     []*NextStep
	Disposal Disposal
}

type NextStep struct {
	Step      string
	Operation string
}
