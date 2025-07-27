package test

import (
	"testing"
	"calculator-devops/internal/calc"
)
func TestDivide(t *testing.T) {
	tests:=[]struct {
		a,b int 
		want int
		wantErr bool
	} {
		{a:8, b:4, want:2, wantErr:false},
		{a:5,b:0, want:0,wantErr:true},
		{a:9,b:3, want:3,wantErr:false},
	}

	for _,tt:=range tests{

		got,err:=calc.Divide(tt.a,tt.b)

		if(tt.wantErr && err==nil){
			t.Errorf("Expected error for input %d/%d, but got none", tt.a,tt.b)
		}

		if(!tt.wantErr && got!=tt.want){
			t.Errorf("Divide (%d,%d)=%d;want %d",tt.a,tt.b,got,tt.want)
		}

	}
}
