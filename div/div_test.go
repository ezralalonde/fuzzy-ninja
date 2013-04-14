package div

import (
	"fmt"
	"math/big"
	"testing"
)

func tSub(inStr, outStr string, t *testing.T) {
	f := Sub11
	var in big.Int
	_, err := fmt.Sscan(inStr, &in)
	var out big.Int
	if err != nil {
		panic("couldn't parse in")
	}
	_, err = fmt.Sscan(outStr, &out)
	if err != nil {
		panic("couldn't parse out")
	}
	mid := f(&in)
	if mid.Cmp(&out) != 0 {
		t.Errorf("Sub11(%v) = %v, wanted %v", in, mid, out)
	}
}

func TestSub11_1(t *testing.T) {
	tSub("12345678901234567900", "1234567890123456790", t)
}
func TestSub11_2(t *testing.T) {
	tSub("1234567890123456790", "123456789012345679", t)
}
func TestSub11_3(t *testing.T) {
	tSub("123456789012345679", "12345678901234558", t)
}

func TestSub11_4(t *testing.T) {
	tSub("12345678901234558", "1234567890123447", t)
}

func TestSub11_5(t *testing.T) {
	tSub("1234567890123447", "123456789012337", t)
}

func TestSub11_6(t *testing.T) {
	tSub("123456789012337", "12345678901226", t)
}

func TestSub11_7(t *testing.T) {
	tSub("12345678901226", "1234567890116", t)
}

func tDig(inStr string, out bool, t *testing.T) {
	f := IsTwoDigits
	var in big.Int
	_, err := fmt.Sscan(inStr, &in)
	if err != nil {
		panic("couldn't parse in")
	}
	ans := f(in)
	if ans != out {
		t.Errorf("IsTwoDigits(%v) = %v, wanted %v", in, ans, out)
	}
}

func TestDig1(t *testing.T) {
	tDig("100", false, t)
}
func TestDig2(t *testing.T) {
	tDig("99", true, t)
}
func TestDig3(t *testing.T) {
	tDig("10", true, t)
}
func TestDig4(t *testing.T) {
	tDig("09", false, t)
}

func tCal(inStr string, out bool, t *testing.T) {
	f := Calc
	var in big.Int
	_, err := fmt.Sscan(inStr, &in)
	if err != nil {
		panic("couldn't parse in")
	}
	ori := in
	ans := f(&in)
	if ans != out {
		t.Errorf("Calc(%v) = %v, wanted %v", ori, ans, out)
	}
}

func TestCalc1(t *testing.T) {
	tCal("11", false, t)
}
func TestCalc2(t *testing.T) {
	tCal("111", true, t)
}
func TestCalc3(t *testing.T) {
	tCal("111111", true, t)
}
func TestCalc4(t *testing.T) {
	tCal("99", false, t)
}
func TestCalc5(t *testing.T) {
	tCal("999", true, t)
}
