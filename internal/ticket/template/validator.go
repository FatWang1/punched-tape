package template

import (
	"errors"
	"fmt"

	"github.com/FatWang1/fatwang-go-utils/desc/set"
	"github.com/FatWang1/flowLite/internal/models"
)

// Validator is the interface that wraps the Validate method.
type Validator interface {
	Validate(models.TicketTemplate) error
}

type validator struct {
	signTypeSet set.Set[string]
}

var (
	ErrStartStepEmpty    = errors.New("start step is empty")
	ErrConfigEmpty       = errors.New("config is empty")
	ErrBadStepConfig     = errors.New("bad step config")
	ErrStepEmpty         = errors.New("step is empty")
	ErrBadSignType       = errors.New("bad sign type")
	ErrNextStepEmpty     = errors.New("next step is empty in non-end step")
	ErrBadNextStep       = errors.New("bad next step")
	ErrEndStepHasNext    = errors.New("end step has next steps")
	ErrDuplicateStep     = errors.New("duplicate step definition")
	ErrStartStepNotFound = errors.New("start step not found in configurations")
	ErrUnreachableSteps  = errors.New("some steps are unreachable")
)

func (v *validator) Validate(tpl models.TicketTemplate) error {
	if len(tpl.StartStep) == 0 {
		return ErrStartStepEmpty
	}
	if len(tpl.Config) == 0 {
		return ErrConfigEmpty
	}

	stepMap := make(map[string]*models.TicketConfig, len(tpl.Config))
	endStepSet := set.Setify(tpl.EndStep...)
	for _, c := range tpl.Config {
		if c == nil || len(c.Step) == 0 {
			return ErrBadStepConfig
		}
		if len(c.Next) == 0 && !endStepSet.HasKey(c.Step) {
			return ErrNextStepEmpty
		}
		if !v.signTypeSet.HasKey(c.Disposal.SignType) {
			return ErrBadSignType
		}
		if len(c.Next) > 0 && endStepSet.HasKey(c.Step) {
			return ErrEndStepHasNext
		}
		if _, exists := stepMap[c.Step]; exists {
			return ErrDuplicateStep
		}
		stepMap[c.Step] = c
	}
	if _, ok := stepMap[tpl.StartStep]; !ok {
		return ErrStartStepNotFound
	}

	// 验证是否存在不可达的步骤
	if err := validateReachability(tpl.StartStep, stepMap, endStepSet); err != nil {
		return err
	}
	return nil
}

func validateReachability(start string, stepMap map[string]*models.TicketConfig, endStepSet set.Set[string]) error {
	visited := set.InitSet[string](len(stepMap))
	queue := make([]string, 0, len(stepMap))
	queue = append(queue, start)

	for len(queue) > 0 {
		currentStep := queue[0]
		queue = queue[1:]
		if visited.HasKey(currentStep) {
			continue
		}
		visited.Set(currentStep)
		config, exists := stepMap[currentStep]
		if !exists {
			return fmt.Errorf("%w: %s", ErrUnreachableSteps, currentStep)
		}
		if endStepSet.HasKey(currentStep) {
			// 结束步骤
			continue
		}
		for _, next := range config.Next {
			if next.GetStep() == "" {
				return ErrBadNextStep
			}
			queue = append(queue, next.Step)
		}
	}

	if visited.Len() != len(stepMap) {
		return ErrUnreachableSteps
	}
	return nil
}
