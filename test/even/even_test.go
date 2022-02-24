package even

import "testing"

func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log(" 10 must be even!")
		t.Fail()
	}
	if Even(7) {
		t.Log(" 7 is not even!")
		t.Fail()
	}

}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}

/*
用（测试数据）表驱动测试
*/
var tests = []struct { // Test table
	in  int
	out bool
}{
	{1, false},
	{2, true},
	{3, false},
}

func verify(t *testing.T, testnum int, testcase string, input int, output, expected bool) {
	if expected != output {
		t.Errorf("%d. %s with input = %s: output %s != %s", testnum, testcase, input, output, expected)
	}
}
func TestFunction(t *testing.T) {
	for i, tt := range tests {
		s := Even(tt.in)
		verify(t, i, "Even:", tt.in, s, tt.out)
	}
}
