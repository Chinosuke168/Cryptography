package main 
import ("fmt")

func xorEncrypt(text string, key byte) string{
		result := make([]byte, len(text))
		for i := 0; i < len (text); i++ {
			result [i] = text [i] ^ key
		}
		return string(result)
}


func main(){
	var plaintext string 
	var key byte

	fmt.Print("Enter your message: ")
	fmt.Scanln(&plaintext)

	fmt.Print("Enter a single-byte key (Character): ")
	fmt.Scanf(" %c", &key)

	ciphertext := xorEncrypt(plaintext, key)
	fmt.Println("\nEncrypted text: ", ciphertext)

	decrypted := xorEncrypt(ciphertext, key)
	fmt.Println("Decrypted text: ", decrypted)
}
