package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)




func main(){
	var p string
	fmt.Print("Enter a string:")
	fmt.Scanln(&p)

	binary := ""
	for i := 0; i < len(p); i++ {
		binary += fmt.Sprintf("%08b", p[i])
	}


	hexadecimal := hex.EncodeToString([]byte(p))

	base64Encoded := base64.StdEncoding.EncodeToString([]byte(p))

	fmt.Println("\n<<<<<<<<<<<<<<< Encoded Results >>>>>>>>>>>>>>>")
	fmt.Println("Plain text: ", p)
	fmt.Println("Binary: ", binary)
	fmt.Println("Hexadecimal: ", hexadecimal)
	fmt.Println("Base64: ", base64Encoded)
}

