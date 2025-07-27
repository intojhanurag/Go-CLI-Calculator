package test

import (
	"testing"
	"calculator-devops/internal/calc"
)
func TestMultiply(t *testing.T) {
	tests:=[]struct {
		a,b int 
		want int
		
	} {
		{a:8, b:4, want:32},
		{a:5,b:0, want:0},
		{a:9,b:-3, want:-27},
	}

	for _,tt:=range tests{

		got:=calc.Multiply(tt.a,tt.b)
		if(got!=tt.want){
			t.Errorf("Multiply (%d,%d)=%d;want %d",tt.a,tt.b,got,tt.want)
		}

	}
}

