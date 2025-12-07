package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	s := []int{1,2,3}
	s1 := s[0:2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)
}