package main

import "fmt"

func info(s shape) float64 {
	return s.dienTich()
}

func main() {
	circle := circle{radius: 2.4}
	square := square{length: 3}

	fmt.Println(info(circle))
	fmt.Println(info(square))

}
