package vec3

import (
	"reflect"
	"testing"
)

func TestDot(t *testing.T) {
	type args struct {
		v1 *Vec3
		v2 *Vec3
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dot(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitVector(t *testing.T) {
	type args struct {
		v *Vec3
	}
	tests := []struct {
		name string
		args args
		want *Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnitVector(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Add(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	type args struct {
		rhs *Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec3{
				e: tt.fields.e,
			}
			if got := v.Add(tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_B(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				e: tt.fields.e,
			}
			if got := v.B(); got != tt.want {
				t.Errorf("B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Div(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	type args struct {
		rhs float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec3{
				e: tt.fields.e,
			}
			if got := v.Div(tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_G(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				e: tt.fields.e,
			}
			if got := v.G(); got != tt.want {
				t.Errorf("G() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Length(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				e: tt.fields.e,
			}
			if got := v.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Mul(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	type args struct {
		rhs float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec3{
				e: tt.fields.e,
			}
			if got := v.Mul(tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_R(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				e: tt.fields.e,
			}
			if got := v.R(); got != tt.want {
				t.Errorf("R() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Sub(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	type args struct {
		rhs *Vec3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vec3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vec3{
				e: tt.fields.e,
			}
			if got := v.Sub(tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_X(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				e: tt.fields.e,
			}
			if got := v.X(); got != tt.want {
				t.Errorf("X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Y(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				e: tt.fields.e,
			}
			if got := v.Y(); got != tt.want {
				t.Errorf("Y() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVec3_Z(t *testing.T) {
	type fields struct {
		e [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vec3{
				e: tt.fields.e,
			}
			if got := v.Z(); got != tt.want {
				t.Errorf("Z() = %v, want %v", got, tt.want)
			}
		})
	}
}
