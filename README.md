# Cypher Tool

## Description
This is a basic command-line tool that allows you to encrypt and decrypt messages using various encryption techniques. The tool supports three different encryption methods:

ROT13: A cypher that shifts each letter by 13 positions.
Reverse Alphabet: A cypher where each letter is replaced by its counterpart in the reversed alphabet.
Caesar Cypher: A classic cypher that shifts each letter by a fixed number of positions (in this case, 3).
How to Use
1. Run the Program
You can run the program using the following command:

go run main.go

The program will guide you through the process step by step:

2. Select an Operation
You’ll first need to choose whether to encrypt or decrypt a message:

Select operation (1/2):
1. Encrypt
2. Decrypt

Type 1 for encrypt or 2 for decrypt, and press Enter.

3. Select a Cypher
Next, you’ll choose the type of encryption to use:

Select cypher (1/3):
1. ROT13
2. Reverse Alphabet
3. Caesar Cypher

Type 1, 2, or 3 based on which encryption you want to use, and press Enter.

4. Enter the Message
After choosing the encryption type, you’ll be prompted to enter the message you want to encrypt or decrypt:

Enter the message:

Type your message and press Enter.

5. See the Result
Once you’ve entered the message, the program will output the result:

Result: <your encrypted or decrypted message>

Encryption Methods
1. ROT13
How it works: Shifts each letter of the message by 13 positions in the alphabet. This means that letters in the first half of the alphabet (A-M) become letters in the second half (N-Z), and vice versa.
Example:
Original: HELLO
ROT13: URYYB
2. Reverse Alphabet
How it works: Each letter is replaced by its counterpart in the reversed alphabet. For example, A becomes Z, B becomes Y, and so on.
Example:
Original: HELLO
Reverse Alphabet: SVOOL
3. Caesar Cypher (Shift of 3)
How it works: This is a classic cipher that shifts each letter by 3 positions in the alphabet. So, A becomes D, B becomes E, etc.
Example:
Original: HELLO
Caesar Cypher (shift 3): KHOOR
Example Usage
Here’s an example of how the tool works:


$ go run main.go

Welcome to the Cypher Tool!

Select operation (1/2):
1. Encrypt
2. Decrypt
1

Select cypher (1/3):
1. ROT13
2. Reverse Alphabet
3. Caesar Cypher
1

Enter the message:
hello

Result:
uryyb
In this example, the user chose to encrypt using the ROT13 cipher, and the message "hello" was encrypted to "uryyb".

Summary
This is a simple encryption and decryption tool built as a learning project. It allows users to experiment with three different encryption techniques (ROT13, Reverse Alphabet, and Caesar Cypher).





