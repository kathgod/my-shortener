package main

import "testing"

func TestMyTest(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyTest(tt.args.n); got != tt.want {
				t.Errorf("MyTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
