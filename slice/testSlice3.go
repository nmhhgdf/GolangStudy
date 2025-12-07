package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	var slice2 []int

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	
	slice2 = make([]int, 3)
	slice2[0] = 100
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	slice4 := make([]int , 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	var slice5 []int
	if (slice5 == nil) {
		fmt.Println("null")
	} else {
		fmt.Println("not null")
	}
}