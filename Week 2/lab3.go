package main
import ("fmt")

func myAND(a, b int){
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
}

func myOR( a, b  int ){
	fmt.Printf("%d | %d = %d \n", a, b, a|b)	
}

func myXOR(a,b int){
	fmt.Printf("%d ^ %d = %d \n", a, b, a^b)
}

func myNOT(a int){
	fmt.Printf("^%d = %d \n", a, ^a)
}

func myShift(a int){
	fmt.Printf("%d << 1 = %d\n", a, a<<1)
	fmt.Printf("%d >> 1 = %d \n", a, a>>1)
}


func main() {
	var a, b int

	fmt.Printf("Input first number:")
	fmt.Scan(&a)
	fmt.Printf("Input second number:")
	fmt.Scan(&b)

	fmt.Println("Bitwise Operations:")
	myAND(a, b)
	myOR(a, b)
	myXOR(a, b)
	myNOT(a)
	myShift(a)

}