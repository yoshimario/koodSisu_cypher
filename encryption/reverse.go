package encryption



// Encrypt the message with reverse
func Encrypt_reverse(s string) string {

	return reverseString(s)
}


// Decrypt the message with reverse
func Decrypt_reverse(s string) string {

	return reverseString(s)

}


func reverseString(s string) string {

	r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
