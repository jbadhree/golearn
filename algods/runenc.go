package algods

import "fmt"

func ExecRunLengthEncoding() {
	// Declare a string
	str := "aaabbbccccdddeeee"
	// Call the run length encoding function
	encoded := ""
	count := 1
	previousChar := str[0]
	strlen := len(str)
	for i := 1; i < strlen; i++ {
		s := string(str[i])
		if s == string(previousChar) {
			count++
		} else {
			encoded += string(previousChar) + fmt.Sprint(count)
			count = 1
			previousChar = str[i]
		}
	}

	encoded += string(previousChar) + fmt.Sprint(count)

	// Print the encoded string
	fmt.Println(encoded)
}
