package main

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]

}

func main() {

	number_comma := comma("213000345678")
	println(number_comma)

}
