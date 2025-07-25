package io

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winkoz/plonk/internal/network"
)

func Test_stitcher_Stitch(t *testing.T) {
	fixturesPath := "../fixtures/stitcher"
	targetPath := "/tmp/plonk/tests/deploy"
	service := service{networkService: network.NewService()}
	service.DeletePath(targetPath)
	service.CreatePath(targetPath)

	type args struct {
		sourcePath     string
		targetPath     string
		targetFilename string
		filePaths      []string
		transformator  Transformator
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantChannel []byte
		want        string
		service     Service
	}{
		{
			name: "succeeds in merging all sources into a single output.yml file",
			args: args{
				sourcePath:     fixturesPath,
				targetPath:     targetPath,
				targetFilename: "output.yml",
				filePaths: []string{
					"1.yaml",
					"2.yaml",
					"3.yaml",
				},
				transformator: NoOpTransformator,
			},
			wantErr:     false,
			wantChannel: nil,
			want: `- file1
- file1_line2
- name: service
- line 1
- line 2
`,
			service: service,
		},
		{
			name: "checks the transformator is used correctly",
			args: args{
				sourcePath:     fixturesPath,
				targetPath:     targetPath,
				targetFilename: "output.yml",
				filePaths: []string{
					"1.yaml",
					"2.yaml",
					"3.yaml",
				},
				transformator: func(input []byte) []byte {
					return []byte("something else")
				},
			},
			wantErr:     false,
			wantChannel: nil,
			want:        "something else",
			service:     service,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stitcher{
				service: tt.service,
			}
			if err := s.Stitch(tt.args.sourcePath, tt.args.targetPath, tt.args.targetFilename, tt.args.filePaths, tt.args.transformator); (err != nil) != tt.wantErr {
				t.Errorf("stitcher.Stitch() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				outputPath := fmt.Sprintf("%s/%s", tt.args.targetPath, tt.args.targetFilename)
				outputBytes, err := service.ReadFile(outputPath)
				if err != nil {
					t.Errorf("stitcher.Stitch() output_read_error = %v", err)
				}
				outputContents := string(outputBytes)
				assert.Equal(t, tt.want, outputContents)
			}
		})
	}
}
