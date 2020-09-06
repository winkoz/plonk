package io

import (
	"reflect"
	"testing"
)

func TestNewInterpolator(t *testing.T) {
	tests := []struct {
		name string
		want Interpolator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInterpolator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReplacer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interpolator_SubstituteValues(t *testing.T) {
	sut := interpolator{}
	type args struct {
		source   map[string]string
		template string
	}
	tests := []struct {
		name    string
		r       interpolator
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "SubstituteValues should replace all source keys on the template with their respective values",
			r:    sut,
			args: args{
				source: map[string]string{
					"var1": "value1",
					"var2": "value2",
				},
				template: "Hi $var1 this var2 should not change; but this other $var1 should!",
			},
			want:    "Hi value1 this var2 should not change; but this other value1 should!",
			wantErr: false,
		},
		{
			name: "SubstituteValues should ONLY replace keys present in the template prefixed with $",
			r:    sut,
			args: args{
				source: map[string]string{
					"var1": "value1",
					"var2": "value2",
				},
				template: "Hi $var1 this is $var2!",
			},
			want:    "Hi value1 this is value2!",
			wantErr: false,
		},
		{
			name: "SubstituteValues should not recursively substitute keys inside the template and source maps",
			r:    sut,
			args: args{
				source: map[string]string{
					"var1": "$var2",
					"var2": "value2",
				},
				template: "Hi $var1 this is $var2!",
			},
			want:    "Hi $var2 this is value2!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := interpolator{}
			got, err := r.SubstituteValues(tt.args.source, tt.args.template)
			if (err != nil) != tt.wantErr {
				t.Errorf("replacer.SubstituteValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("replacer.SubstituteValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
