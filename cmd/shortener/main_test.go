package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Тестируем функции из мейна
func TestMyTest(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, myTest(tt.args.n), tt.want, "MyTest() = %v, want %v")
		})
	}
}
