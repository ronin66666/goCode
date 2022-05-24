package main

import "fmt"

func main() {
	s := []int { 1, 2, 3, 4 }
	fmt.Printf("&s = %p", s)
	fmt.Printf("&s = %p", &s)

	s = append(s, 5)
	fmt.Println(s)
	fmt.Printf("&s = %p", s)
	fmt.Printf("&s = %p", &s)
}
