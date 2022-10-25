package main

import "fmt"

func main() {
	// Dealing with strings
	var firstName string = "Wei-Meng"

	address := "The White House\n1600 Pennsylvania Avenue NW\nWashington, DC 20500"

	//Strings can contain special characters.

	//There is also raw strings special characters have no meaning.span multiple lines.
	//raw strings use back ticks.

	quotation := `"Anyone who has never made
a mistake has never tried
anything new."
--Albert Einstein`

	// represent Unicode characters in your strings
	str1 := "你好,世界"   // Chinese
	str2 := "こんにちは世界" // Japanese

	// Unicode encodings of each character in your strings:
	str3 := "\u4f60\u597d\uff0c\u4e16\u754c" // 你好,世界

	fmt.Println(firstName, address, quotation, str1, str2, str3)
}
