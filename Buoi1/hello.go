package main

import "fmt"

func main() {
	var n int
	fmt.Printf("nhap n: ")
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for pos, _ := range arr {
		fmt.Printf("nhap a[%d]: ", pos)
		fmt.Scan(&arr[pos])
	}
	bubbleSort(arr)
	for pos, number := range arr {
		fmt.Printf("\n kết quả a[%d]: %d \n", pos, number)
	}

}

func bubbleSort(sl []int) {
	n := len(sl)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sl[j] > sl[j+1] {
				// Hoán đổi hai phần tử
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
}
