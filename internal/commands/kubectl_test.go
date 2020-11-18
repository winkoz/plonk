package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

func Test_kubecltCommand_Deploy(t *testing.T) {
	executorMock := new(sharedtesting.ExecutorMock)
	ctx := config.Context{}
	type fields struct {
		executor *sharedtesting.ExecutorMock
	}
	type args struct {
		env          string
		manifestPath string
		ctx          config.Context
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantCommand string
		wantErr     bool
	}{
		{
			name: "Successfully calls Deploy with the passed in command",
			fields: fields{
				executor: executorMock,
			},
			args: args{
				env:          "production",
				manifestPath: "this/is/not/a/real/path",
				ctx:          ctx,
			},
			wantCommand: "kubectl -f this/is/not/a/real/path",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := kubecltCommand{
				executor: tt.fields.executor,
			}

			tt.fields.executor.On(
				"Run",
				tt.wantCommand,
			).Return(
				nil,
			)

			err := k.Deploy(tt.args.env, tt.args.manifestPath, tt.args.ctx)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
