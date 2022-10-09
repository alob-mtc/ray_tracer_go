package ray

import (
	"github.com/alob-mtc/ray_tracer/vec3"
	"reflect"
	"testing"
)

func TestRay_Direction(t *testing.T) {
	type fields struct {
		A vec3.Vec3
		B vec3.Vec3
	}
	tests := []struct {
		name   string
		fields fields
		want   vec3.Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := r.Direction(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Direction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRay_Origin(t *testing.T) {
	type fields struct {
		A vec3.Vec3
		B vec3.Vec3
	}
	tests := []struct {
		name   string
		fields fields
		want   vec3.Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := r.Origin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Origin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRay_PointOfParameter(t *testing.T) {
	type fields struct {
		A vec3.Vec3
		B vec3.Vec3
	}
	type args struct {
		t float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *vec3.Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := r.PointOfParameter(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PointOfParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
