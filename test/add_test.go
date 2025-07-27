package test

import (
	"testing"
	"calculator-devops/internal/calc"
)
func TestAdd(t *testing.T) {
	tests:=[]struct {
		a,b int 
		want int
	} {
		{a:8, b:4, want:12},
		{a:5,b:0, want:5},
		{a:9,b:-3, want:6},
	}

	for _,tt:=range tests{

		got:=calc.Add(tt.a,tt.b)
		if(got!=tt.want){
			t.Errorf("Add (%d,%d)=%d;want %d",tt.a,tt.b,got,tt.want)
		}

	}
}
