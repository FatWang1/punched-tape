package template

import (
	"testing"

	"github.com/FatWang1/fatwang-go-utils/desc/set"
	"github.com/FatWang1/flowLite/internal/models"
)

func Test_canTraverse(t *testing.T) {
	type args struct {
		start      string
		stepMap    map[string]*models.TicketConfig
		endStepSet set.Set[string]
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "all is ok",
			args: args{
				start: "1",
				stepMap: map[string]*models.TicketConfig{
					"1": {
						Step: "1",
						Next: []*models.NextStep{
							{
								Step: "2",
							},
							{
								Step: "3",
							},
						},
					},
					"2": {
						Step: "2",
						Next: []*models.NextStep{
							{
								Step: "3",
							},
						},
					},
					"3": {
						Step: "3",
						Next: []*models.NextStep{
							{
								Step: "1",
							},
							{
								Step: "4",
							},
						},
					},
					"4": {
						Step: "4",
					},
				},
				endStepSet: set.Setify("4"),
			},
			wantErr: false,
		},
		{
			name:    "",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateReachability(tt.args.start, tt.args.stepMap, tt.args.endStepSet); (err != nil) != tt.wantErr {
				t.Errorf("canTraverse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
