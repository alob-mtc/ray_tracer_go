package vec3

import (
	"reflect"
	"testing"
)

func TestVec3_Add(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	type args struct {
		newV *Vec3
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec3{
				e: tt.fields.e,
			}
			if got := v.Add(tt.args.newV); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
