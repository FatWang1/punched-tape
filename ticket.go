package punched_tape

import (
	"github.com/FatWang1/punched-tape/models"
)

type TicketBuilder struct {
	option models.Ticket
}

func NewTicketBuilder() *TicketBuilder {
	return &TicketBuilder{}
}

func (b *TicketBuilder) SetUid(uid string) *TicketBuilder {
	b.option.Uid = uid
	return b
}

func (b *TicketBuilder) SetOrderNum(orderNum string) *TicketBuilder {
	b.option.OrderNum = orderNum
	return b
}

func (b *TicketBuilder) SetMemo(memo string) *TicketBuilder {
	b.option.Memo = memo
	return b
}

func (b *TicketBuilder) Build() models.Ticket {
	return b.option
}
