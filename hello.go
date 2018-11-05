package main

import (
	"fmt"
)

func main() {
	a := 5
	ptr := &a

	slInt := []int{3, 324, 234, 6, 1, 43, 76, 23, 37}
	arInt := [3]int{3, 4, 66}
	var ptrI *[]int = &slInt
	ptrI2 := &slInt

	fmt.Printf(fmt.Sprintf("%-7v", "ptr")+" è %T ha indirizzo %v e valore %v \n", ptr, &ptr, *ptr)
	fmt.Printf(fmt.Sprintf("%-7v", "arInt")+" è %T ha indirizzo %v e valore %v \n", arInt, &arInt, arInt)
	fmt.Printf(fmt.Sprintf("%-7v", "slInt")+" è %T ha indirizzo %v e valore %v \n", slInt, &slInt, slInt)
	fmt.Printf(fmt.Sprintf("%-7v", "ptrI")+" è %T ha indirizzo %v e valore %v \n", ptrI, &ptrI, *ptrI)
	fmt.Printf(fmt.Sprintf("%-7v", "ptrI2")+" è %T ha indirizzo %v e valore %v \n", ptrI2, &ptrI2, *ptrI2)

}
