package punched_tape

import "github.com/FatWang1/punched-tape/internal/models"

type TemplateBuilder struct {
	option models.TicketTemplate
}

func NewTemplateBuilder() *TemplateBuilder {
	return &TemplateBuilder{}
}

func (b *TemplateBuilder) SetBuiltin(builtin bool) *TemplateBuilder {
	b.option.Builtin = builtin
	return b
}

func (b *TemplateBuilder) SetUid(uid string) *TemplateBuilder {
	b.option.Uid = uid
	return b
}

func (b *TemplateBuilder) SetStartStep(startStep string) *TemplateBuilder {
	b.option.StartStep = startStep
	return b
}

func (b *TemplateBuilder) SetEndStep(endStep ...string) *TemplateBuilder {
	b.option.EndStep = endStep
	return b
}

func (b *TemplateBuilder) AddConfig(config ...*models.StepConfig) *TemplateBuilder {
	b.option.Config = append(b.option.Config, config...)
	return b
}

func (b *TemplateBuilder) Build() models.TicketTemplate {
	return b.option
}
