package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Get the input data required for the operation
func GetInput() (toEncrypt bool, encoding string, message string) {

	var invalidInput bool
	reader := bufio.NewReader (os.Stdin)

	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println("")
	fmt.Println("Select operation (1/2)")
	fmt.Println("1. Encrypt.")
	fmt.Println("2. Decrypt.")






	fmt.Println("")
	fmt.Println("Select cypher (1/2)")
	fmt.Println("1. ROT13.")
	fmt.Println("2. Reverse.")
	fmt.Println("")
	fmt.Println("Enter the message:")

	if invalidInput {


	}



}