package main

import (
	"fmt"
	"fuzzy-ninja/div"
	"math/big"
)

func main() {
	var lines int
	_, err := fmt.Scanf("%d", &lines)
	if err != nil {
		panic("can't read number of lines")
	}

	for ii := 0; ii < lines; ii++ {
		var val big.Int
		_, err := fmt.Scanf("%v", &val)
		if err != nil {
			panic("can't read line")
		}
		old := val
		for more := true; more; more = div.Calc(&val) {
			fmt.Println(val.String())
		}
		if div.IsDivisibleBy11(old) {
			fmt.Printf("The number %s is divisible by 11.\n\n", old.String())
		} else {
			fmt.Printf("The number %s is not divisible by 11.\n\n", old.String())
		}
	}
}
