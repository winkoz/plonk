package io

import (
	"fmt"
	"os"
	"testing"

	"github.com/winkoz/plonk/internal/sharedtesting"
)

func Test_duplicator_CopyMultiple(t *testing.T) {
	fixturesPath := "../fixtures/scripts"
	testTargetPath := "/tmp/plonk/tests/deploy"
	sharedtesting.CreatePath(testTargetPath)
	type args struct {
		sourcePath    string
		targetPath    string
		sources       []string
		transformator Transformator
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "succeeds by copying all sources from one path to the other",
			args: args{
				sourcePath: fixturesPath,
				targetPath: testTargetPath,
				sources: []string{
					"service.yaml",
				},
				transformator: sharedtesting.SimpleTransformator,
			},
			wantErr: false,
		},
		{
			name: "throws error when the sourcePath doesn't exist",
			args: args{
				sourcePath: "../fixtures/missingscriptsfolder",
				targetPath: testTargetPath,
				sources: []string{
					"service.yaml",
				},
			},
			wantErr: true,
		},
		{
			name: "throws error when the targetPath doesn't exist",
			args: args{
				sourcePath: fixturesPath,
				targetPath: "/tmp/plonk/thisfolderdoesntexist",
				sources: []string{
					"service.yaml",
				},
			},
			wantErr: true,
		},
		{
			name: "throws error when one of the source files doesn't exist",
			args: args{
				sourcePath: fixturesPath,
				targetPath: testTargetPath,
				sources: []string{
					"service.yaml",
					"thisdoesntexist.yaml",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := duplicator{}
			if err := d.CopyMultiple(tt.args.sourcePath, tt.args.targetPath, tt.args.sources, tt.args.transformator); (err != nil) != tt.wantErr {
				t.Errorf("duplicator.CopyMultiple() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				for _, s := range tt.args.sources {
					targetSourcePath := fmt.Sprintf("%s/%s", tt.args.targetPath, s)
					if _, err := os.Stat(targetSourcePath); (err != nil) != tt.wantErr {
						t.Errorf("duplicator.CopyMultiple() test_check_error = %v", err)
					}
				}
			}
		})
	}
	sharedtesting.DeletePath(testTargetPath)
}
