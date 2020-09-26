package io

import (
	"fmt"
	"os"
	"testing"
)

func Test_duplicator_CopyMultiple(t *testing.T) {
	fixturesPath := "../fixtures/scripts"
	testTargetPath := "/tmp/plonk/tests/deploy"
	CreatePath(testTargetPath)
	type args struct {
		targetPath      string
		sourceLocations []FileLocation
		transformator   Transformator
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
				sourceLocations: []FileLocation{
					{
						OriginalFilePath: "/service.yaml",
						ResolvedFilePath: fixturesPath + "/service.yaml",
					},
				},
				transformator: NoOpTransformator,
			},
			wantErr: false,
		},
		{
			name: "throws error when the sourcePath doesn't exist",
			args: args{
				targetPath: testTargetPath,
				sourceLocations: []FileLocation{
					{
						OriginalFilePath: "/service.yaml",
						ResolvedFilePath: "../fixtures/missingscriptsfolder" + "/service.yaml",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "throws error when the targetPath doesn't exist",
			args: args{
				targetPath: "/tmp/plonk/thisfolderdoesntexist",
				sourceLocations: []FileLocation{
					{
						OriginalFilePath: "/service.yaml",
						ResolvedFilePath: fixturesPath + "/service.yaml",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "throws error when one of the source files doesn't exist",
			args: args{
				targetPath: testTargetPath,
				sourceLocations: []FileLocation{
					{
						OriginalFilePath: "/service.yaml",
						ResolvedFilePath: fixturesPath + "/service.yaml",
					},
					{
						OriginalFilePath: "/thisdoesntexist.yaml",
						ResolvedFilePath: fixturesPath + "/thisdoesntexist.yaml",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := duplicator{}
			if err := d.CopyMultiple(tt.args.targetPath, tt.args.sourceLocations, tt.args.transformator); (err != nil) != tt.wantErr {
				t.Errorf("duplicator.CopyMultiple() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				for _, s := range tt.args.sourceLocations {
					targetSourcePath := fmt.Sprintf("%s/%s", tt.args.targetPath, s.OriginalFilePath)
					if _, err := os.Stat(targetSourcePath); (err != nil) != tt.wantErr {
						t.Errorf("duplicator.CopyMultiple() test_check_error = %v", err)
					}
				}
			}
		})
	}
	DeletePath(testTargetPath)
}
