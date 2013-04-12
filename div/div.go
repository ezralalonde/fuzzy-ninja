package div

import (
	"fmt"
	"math/big"
)

func Sub11(xx big.Int) big.Int {
	mm := new(big.Int)
	kk := big.NewInt(10)
	yy := new(big.Int)
	yy.DivMod(&xx, kk, mm)
	yy.Sub(yy, mm)
	return *yy
}

func IsTwoDigits(xx big.Int) bool {
	return xx.Cmp(big.NewInt(100)) == -1 && xx.Cmp(big.NewInt(9)) == 1
}

func Calc(xx big.Int, loud bool) bool {
	var yy big.Int
	old := xx
	for yy = Sub11(xx); xx.Cmp(big.NewInt(9)) == 1; xx, yy = yy, Sub11(yy) {
		if loud {
			fmt.Printf("%s\n", xx.String())
		}
	}
	if xx.Mod(&xx, big.NewInt(11)).Cmp(big.NewInt(0)) == 0 {
		if loud {
			fmt.Printf("The number %s is divisible by 11.\n\n", old.String())
		}
		return true
	}
	if loud {
		fmt.Printf("The number %s is not divisible by 11.\n\n", old.String())
	}
	return false
}
