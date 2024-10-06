package main

import (
    "testing"
)

// TestEncryptRot13 tests the encrypt_rot13 function
func TestEncryptRot13(t *testing.T) {
    input := "hello"
    expected := "uryyb" // ROT13 of "hello"
    result := encrypt_rot13(input)

    if result != expected {
        t.Errorf("encrypt_rot13 failed, expected: %s, got: %s", expected, result)
    }
}

// TestDecryptRot13 tests the decrypt_rot13 function
func TestDecryptRot13(t *testing.T) {
    input := "uryyb"
    expected := "hello" // ROT13 of "uryyb" is "hello"
    result := decrypt_rot13(input)

    if result != expected {
        t.Errorf("decrypt_rot13 failed, expected: %s, got: %s", expected, result)
    }
}

// TestEncryptReverse tests the encrypt_reverse function
func TestEncryptReverse(t *testing.T) {
    input := "abc"
    expected := "zyx" // Reverse of "abc" is "zyx"
    result := encrypt_reverse(input)

    if result != expected {
        t.Errorf("encrypt_reverse failed, expected: %s, got: %s", expected, result)
    }
}

// TestDecryptReverse tests the decrypt_reverse function
func TestDecryptReverse(t *testing.T) {
    input := "zyx"
    expected := "abc" // Reverse of "zyx" is "abc"
    result := decrypt_reverse(input)

    if result != expected {
        t.Errorf("decrypt_reverse failed, expected: %s, got: %s", expected, result)
    }
}

// TestEncryptCaesar tests the encrypt_caesar function
func TestEncryptCaesar(t *testing.T) {
    input := "hello"
    shift := 3
    expected := "khoor" // Caesar shift of 3 on "hello" gives "khoor"
    result := encrypt_caesar(input, shift)

    if result != expected {
        t.Errorf("encrypt_caesar failed, expected: %s, got: %s", expected, result)
    }
}

// TestDecryptCaesar tests the decrypt_caesar function
func TestDecryptCaesar(t *testing.T) {
    input := "khoor"
    shift := 3
    expected := "hello" // Decrypting "khoor" with shift 3 gives "hello"
    result := decrypt_caesar(input, shift)

    if result != expected {
        t.Errorf("decrypt_caesar failed, expected: %s, got: %s", expected, result)
    }
}
