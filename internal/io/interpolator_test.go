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
		// TODO: https://github.com/winkoz/plonk/issues/58 Add test cases.
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
	sut := NewInterpolator()
	type args struct {
		source   map[string]string
		template string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "SubstituteValues should replace all source keys on the template with their respective values",
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
			got := sut.SubstituteValues(tt.args.source, tt.args.template)
			if got != tt.want {
				t.Errorf("interpolator.SubstituteValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interpolator_SubstituteValuesInMap(t *testing.T) {
	source := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	type args struct {
		source map[string]string
		target map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "successfully replaces all interpolatable keys from the values of the target map",
			args: args{
				source: source,
				target: map[string]string{
					"interpolated":     "this-is-$key1",
					"non-interpolated": "just-a-value",
					"interpolated2":    "another-$key3-interpolated-value",
				},
			},
			want: map[string]string{
				"interpolated":     "this-is-value1",
				"non-interpolated": "just-a-value",
				"interpolated2":    "another-value3-interpolated-value",
			},
			wantErr: false,
		},
		{
			name: "successfully returns the same target map when there are no interpolatable values",
			args: args{
				source: source,
				target: map[string]string{
					"non-interpolated":  "this-is-a-value",
					"non-interpolated2": "just-a-value",
					"non-interpolated3": "another-non-interpolated-value",
				},
			},
			want: map[string]string{
				"non-interpolated":  "this-is-a-value",
				"non-interpolated2": "just-a-value",
				"non-interpolated3": "another-non-interpolated-value",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := interpolator{}
			got := i.SubstituteValuesInMap(tt.args.source, tt.args.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("interpolator.SubstituteValuesInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
