package encryption

// Encrypt the message with rot13
func Encrypt_rot13(s string) string {

	return rot13String(s)
}




// Decrypt the message with rot13
func Decrypt_rot13(s string) string {

	return rot13String(s)
}

func rot13String(s string) string{

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