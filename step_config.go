package punched_tape

import (
	"github.com/FatWang1/punched-tape/models"
)

type StepConfigBuilder struct {
	option *models.StepConfig
}

func NewStepConfigBuilder() *StepConfigBuilder {
	return &StepConfigBuilder{option: &models.StepConfig{}}
}

func (b *StepConfigBuilder) SetStep(step string) *StepConfigBuilder {
	b.option.Step = step
	return b
}

func (b *StepConfigBuilder) SetOperator(operator []string) *StepConfigBuilder {
	b.option.Operator = operator
	return b
}

func (b *StepConfigBuilder) SetState(state string) *StepConfigBuilder {
	b.option.State = state
	return b
}

func (b *StepConfigBuilder) SetNext(next []*models.NextStep) *StepConfigBuilder {
	b.option.Next = next
	return b
}

func (b *StepConfigBuilder) SetDisposal(disposal models.Disposal) *StepConfigBuilder {
	b.option.Disposal = disposal
	return b
}

func (b *StepConfigBuilder) Build() *models.StepConfig {
	return b.option
}
