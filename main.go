package main

import (
    "fmt"
    "math/big"
    "fuzzy-ninja/div"
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
        div.Calc(val, true)
    }
}
