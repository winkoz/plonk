package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winkoz/plonk/internal/config"
	"github.com/winkoz/plonk/internal/io"
	"github.com/winkoz/plonk/internal/sharedtesting"
)

func Test_kubecltCommand_Deploy(t *testing.T) {
	executorMock := new(sharedtesting.ExecutorMock)
	ctx := config.Context{
		DeployCommand: "notKubeCtl",
		TargetPath:    "",
	}
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
		wantArgs    []string
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
			wantCommand: "notKubeCtl",
			wantArgs:    []string{"apply", "-f", "this/is/not/a/real/path"},
			wantErr:     false,
		},
		{
			name: "Successfully interpolates the path into the command",
			fields: fields{
				executor: executorMock,
			},
			args: args{
				env:          "production",
				manifestPath: "this/is/not/a/real/path",
				ctx: config.Context{
					DeployCommand: "notKubeCtl -p $PWD",
					TargetPath:    "/this/is/some/path",
				},
			},
			wantCommand: "notKubeCtl",
			wantArgs:    []string{"-p", "/this/is/some/path", "apply", "-f", "this/is/not/a/real/path"},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := kubectlCommand{
				executor:     tt.fields.executor,
				interpolator: io.NewInterpolator(),
				ctx:          tt.args.ctx,
			}

			tt.fields.executor.On(
				"Run",
				tt.wantCommand,
				tt.wantArgs,
			).Return(
				make([]byte, 0), nil,
			)

			err := k.Deploy(tt.args.env, tt.args.manifestPath)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func Test_kubecltCommand_Diff(t *testing.T) {
	executorMock := new(sharedtesting.ExecutorMock)
	ctx := config.Context{
		DeployCommand: "notKubeCtl",
		TargetPath:    "",
	}
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
		wantArgs    []string
		wantErr     bool
	}{
		{
			name: "Successfully calls Diff with the passed in command",
			fields: fields{
				executor: executorMock,
			},
			args: args{
				env:          "production",
				manifestPath: "this/is/not/a/real/path",
				ctx:          ctx,
			},
			wantCommand: "notKubeCtl",
			wantArgs:    []string{"diff", "-f", "this/is/not/a/real/path"},
			wantErr:     false,
		},
		{
			name: "Successfully interpolates the path into the command",
			fields: fields{
				executor: executorMock,
			},
			args: args{
				env:          "production",
				manifestPath: "this/is/not/a/real/path",
				ctx: config.Context{
					DeployCommand: "notKubeCtl -p $PWD",
					TargetPath:    "/this/is/some/path",
				},
			},
			wantCommand: "notKubeCtl",
			wantArgs:    []string{"-p", "/this/is/some/path", "diff", "-f", "this/is/not/a/real/path"},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := kubectlCommand{
				executor:     tt.fields.executor,
				interpolator: io.NewInterpolator(),
				ctx:          tt.args.ctx,
			}

			tt.fields.executor.On(
				"Run",
				tt.wantCommand,
				tt.wantArgs,
			).Return(
				nil,
			)

			err := k.Diff(tt.args.env, tt.args.manifestPath)

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
