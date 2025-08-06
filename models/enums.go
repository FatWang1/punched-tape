package models

import "github.com/FatWang1/fatwang-go-utils/desc/set"

var (
	TicketStatus     = set.Setify(Running, Passed, Rejected)
	DisposalSignType = set.Setify(JointlySign, SerialSign, AnyoneSign)
)
