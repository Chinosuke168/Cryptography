package main
import ("fmt")


func add(a, b float64) float64{
	return a + b
}

func sub(a, b float64) float64{
	return a - b
}

func mul(a, b float64 ) float64{
	return a * b
}

func div(a, b float64) float64{
	if b == 0 {
		fmt.Errorf("division by zero") 
		return 0
	}
	return a / b 
}

func mod(a, b int) int {
	if b == 0 {
		fmt.Errorf("modulo by zero")
		return 0
	}
	return a % b 
}

func main(){
	for {
		fmt.Println("\t <<<<<<<<<< MINI CALCULATOR >>>>>>>>>>>\n\t")
		fmt.Println(" 1.ADD 2.SUB 3.MUL 4.DIV 5.MOD 99.EXIT")
		fmt.Print("Choose: ")

		var choice int 
		fmt.Scan(&choice)

		if choice ==  99 {
			fmt.Println(" \n\tSee you again!! ")
			break
		}

		var a, b float64
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)
		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Println("Result:", add(a,b))
		case 2:
			fmt.Println("Result:", sub(a,b))
		case 3: 
			fmt.Println("Result:", mul(a,b))
		case 4:
			fmt.Println("Result:", div(a,b))
		case 5:
			fmt.Println("Result:", mod(int(a), int(b)))
		default:
			fmt.Println("\nInvalid Choice! please try again.\n")
		}
		fmt.Println()
	}
}