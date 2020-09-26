package io

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test_duplicator_CopyMultiple(t *testing.T) {
	fixturesPath := "../fixtures/scripts"
	testTargetPath := "/tmp/plonk/tests/deploy"
	CreatePath(testTargetPath)
	type args struct {
		targetPath    string
		sourcePaths   []string
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
				targetPath: testTargetPath,
				sourcePaths: []string{
					fixturesPath + "/service.yaml",
				},
				transformator: NoOpTransformator,
			},
			wantErr: false,
		},
		{
			name: "throws error when the sourcePath doesn't exist",
			args: args{
				targetPath: testTargetPath,
				sourcePaths: []string{
					"../fixtures/missingscriptsfolder" + "/service.yaml",
				},
			},
			wantErr: true,
		},
		{
			name: "throws error when the targetPath doesn't exist",
			args: args{
				targetPath: "/tmp/plonk/thisfolderdoesntexist",
				sourcePaths: []string{
					fixturesPath + "/service.yaml",
				},
			},
			wantErr: true,
		},
		{
			name: "throws error when one of the source files doesn't exist",
			args: args{
				targetPath: testTargetPath,
				sourcePaths: []string{
					fixturesPath + "/service.yaml",
					fixturesPath + "/thisdoesntexist.yaml",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := duplicator{}
			if err := d.CopyMultiple(tt.args.targetPath, tt.args.sourcePaths, tt.args.transformator); (err != nil) != tt.wantErr {
				t.Errorf("duplicator.CopyMultiple() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				for _, s := range tt.args.sourcePaths {
					targetSourcePath := fmt.Sprintf("%s/%s", tt.args.targetPath, filepath.Base(s))
					if _, err := os.Stat(targetSourcePath); (err != nil) != tt.wantErr {
						t.Errorf("duplicator.CopyMultiple() test_check_error = %v", err)
					}
				}
			}
		})
	}
	DeletePath(testTargetPath)
}
