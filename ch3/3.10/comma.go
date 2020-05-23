package main

// func comma(s string) string {
// 	n := len(s)
// 	if n <= 3 {
// 		return s
// 	}

// 	return comma(s[:n-3]) + "," + s[n-3:]

// }

func comma(s string) string {
	original_byte_string := []byte(s)
	var new_byte_string []byte

	var count = 0

	for _, element := range original_byte_string {
		count = count + 1
		new_byte_string = append(new_byte_string, element)
		if count == 3 {
			new_byte_string = append(new_byte_string, 44)
			count = 0
		}
	}
	new_string := string(new_byte_string)

	return new_string
}

func main() {

	number_comma := comma("1234567")
	println(number_comma)
}
