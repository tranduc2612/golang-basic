package main

import "fmt"

// // variant: use of slice of byte and conversions
func reverse(s string) string {
	n := len(s)
	c := []byte(s)
	for i := 0; i < n/2; i++ {
		c[i], c[n-1-i] = c[n-1-i], c[i]
	}
	return string(c)
}

func main() {
	var str string
	fmt.Printf("Nhập chuỗi: ")
	fmt.Scan(&str)
	fmt.Print(reverse(str))
}
