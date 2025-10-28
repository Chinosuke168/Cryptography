package main

import ("fmt")

func main(){

	var a, b int
	fmt.Print("Input the first number:")
	fmt.Scan(&a)
	fmt.Print("\n")
	fmt.Print("Input the second number:")
	fmt.Scan(&b)
	fmt.Print("\n")

	bothPositive := (a>0) && (b>0)
	oneGreater := (a>b) || (b>a)
	notEqual := !(a == b)

	fmt.Print("Result: \n")
	fmt.Println("\tBoth are positive: " , bothPositive)
	fmt.Println("\tOne is greater than : ", oneGreater)
	fmt.Println("\tBoth are not equal: ", notEqual)
}

