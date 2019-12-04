package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    min := 146810
    max := 612564
    sum := 0

    for i:=min; i <= max; i++ {
        checkCriteria(i, &sum)
    }

    fmt.Println(sum)
}

func checkCriteria(num int, sum *int) {
    doubleStrFound := false
    neverDecreases := true
    
    numstr := strconv.Itoa(num);
    for i:=0; i < 10; i++ {
        splitstr := strconv.Itoa(i) + strconv.Itoa(i)
        count := strings.SplitAfter(numstr, splitstr)
        if (len(count) > 1) {
            doubleStrFound = true
            break
        }
    }

    numarray := strings.Split(numstr, "")
    for i:=1; i < len(numarray); i++ {
        number1, _ := strconv.Atoi(numarray[i-1])
        number2, _ := strconv.Atoi(numarray[i])
        if number1 > number2 {
            neverDecreases = false
            break
        }
    }

    if doubleStrFound && neverDecreases {
        *sum++
    }
}
