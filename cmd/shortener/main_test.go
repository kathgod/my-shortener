package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_randSeq(t *testing.T) {
	type want struct {
		ln  int
		res string
	}
	tests := []struct {
		name string
		arg  int
		want want
	}{
		{ //TODO: Add test cases.
			name: "test1 Length of return value",
			arg:  6,
			want: want{
				ln:  6,
				res: "123456",
			},
		},
		{
			name: "Test 2 type of return value",
			arg:  5,
			want: want{
				ln:  5,
				res: "12345",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1 := len(randSeq(tt.arg)); got1 != tt.want.ln {
				t.Errorf("randSeq() = %v, want %v", got1, tt.want.ln)
				fmt.Println("Tes1 Fail")
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got2 := reflect.TypeOf(randSeq(tt.arg)); got2 != reflect.TypeOf(tt.want.res) {
				t.Errorf("randSeq() = %v, want %v", got2, tt.want.res)
				fmt.Println("Tes2 Fail")
			}
		})
	}
}
