package test

import (
	"testing"
	"calculator-devops/internal/calc"
)
func TestSubtract(t *testing.T) {
	tests:=[]struct {
		a,b int 
		want int
	} {
		{a:8, b:4, want:4},
		{a:5,b:0, want:5},
		{a:9,b:-3, want:12},
	}

	for _,tt:=range tests{

		got:=calc.Subtract(tt.a,tt.b)
		if(got!=tt.want){
			t.Errorf("Subtract (%d,%d)=%d;want %d",tt.a,tt.b,got,tt.want)
		}

	}
}
