package main

import "fmt"

func printArray(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	// myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4}
	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray = []int{1,2,3,4,5}
	printArray(myArray)

}