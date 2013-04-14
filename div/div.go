package div

import (
	"math/big"
)

func Sub11(xx *big.Int) big.Int {
	mm := new(big.Int)
	kk := big.NewInt(10)
	yy := new(big.Int)
	yy.DivMod(xx, kk, mm)
	yy.Sub(yy, mm)
	return *yy
}

func IsTwoDigits(xx big.Int) bool {
	return xx.Cmp(big.NewInt(100)) == -1 && xx.Cmp(big.NewInt(9)) == 1
}

func Calc(xx *big.Int) bool {
	if IsTwoDigits(*xx) {
		return false
	} else {
		*xx = Sub11(xx)
		return true
	}
	return false
}

func IsDivisibleBy11(xx big.Int) bool {
	return big.NewInt(0).Mod(&xx, big.NewInt(11)).Cmp(big.NewInt(0)) == 0
}
