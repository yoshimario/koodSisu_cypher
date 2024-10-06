package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Greet the user
func greetUser() {
    fmt.Println("Welcome to the Cypher Tool!")
}

// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {
    reader := bufio.NewReader(os.Stdin)

    // Get operation (Encrypt or Decrypt)
    fmt.Println("Select operation (1/2):\n1. Encrypt\n2. Decrypt")
    operation, _ := reader.ReadString('\n')
    operation = strings.TrimSpace(operation)

    if operation == "1" {
        toEncrypt = true
    } else if operation == "2" {
        toEncrypt = false
    } else {
        fmt.Println("Invalid input. Please enter 1 or 2.")
        return getInput() // retry if invalid input
    }

    // Get encoding (ROT13, Reverse, Custom)
    fmt.Println("Select cypher (1/3):\n1. ROT13\n2. Reverse\n3. Custom")
    encoding, _ = reader.ReadString('\n')
    encoding = strings.TrimSpace(encoding)

    if encoding != "1" && encoding != "2" && encoding != "3" {
        fmt.Println("Invalid input. Please enter 1, 2, or 3.")
        return getInput() // retry if invalid input
    }

    // Get message to encrypt/decrypt
    fmt.Println("Enter the message:")
    message, _ = reader.ReadString('\n')
    message = strings.TrimSpace(message)

    return toEncrypt, encoding, message
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {
    return rot13(s)
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {
    return rot13(s)
}

// ROT13 logic
func rot13(s string) string {
    result := ""
    for _, r := range s {
        if r >= 'A' && r <= 'Z' {
            result += string('A' + (r-'A'+13)%26)
        } else if r >= 'a' && r <= 'z' {
            result += string('a' + (r-'a'+13)%26)
        } else {
            result += string(r)
        }
    }
    return result
}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {
    return reverseAlphabet(s)
}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {
    return reverseAlphabet(s)
}

// Reverse Alphabet logic
func reverseAlphabet(s string) string {
    result := ""
    for _, r := range s {
        if r >= 'A' && r <= 'Z' {
            result += string('Z' - (r - 'A'))
        } else if r >= 'a' && r <= 'z' {
            result += string('z' - (r - 'a'))
        } else {
            result += string(r)
        }
    }
    return result
}

// Main logic, invoking other functions
func main() {
    greetUser()
    toEncrypt, encoding, message := getInput()

    var result string
    switch encoding {
    case "1": // ROT13
        if toEncrypt {
            result = encrypt_rot13(message)
        } else {
            result = decrypt_rot13(message)
        }
    case "2": // Reverse Alphabet
        if toEncrypt {
            result = encrypt_reverse(message)
        } else {
            result = decrypt_reverse(message)
        }
    case "3": // Custom encryption (add logic here if needed)
        fmt.Println("Custom encryption not implemented yet.")
    }

    fmt.Println("Result:", result)
}