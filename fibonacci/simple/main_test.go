package main

import "testing"

func TestFiboImpl_Get(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		f    *FiboImpl
		args args
		want int
	}{
		{
			name : "edge case",
			f: &FiboImpl{},
			args: args{
				n: 2,
			},
			want:1,
		},
		{
			name : "edge case",
			f: &FiboImpl{},
			args: args{
				n: 3,
			},
			want:2,
		},
		{
			name : "edge case",
			f: &FiboImpl{},
			args: args{
				n: 7,
			},
			want:13,
		},				
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FiboImpl{}
			if got := f.Get(tt.args.n); got != tt.want {
				t.Errorf("FiboImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
