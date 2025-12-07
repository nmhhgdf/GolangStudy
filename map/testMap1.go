package main

import "fmt"

func main() {
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("null")
	}

	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)
	fmt.Println(len(myMap2))

	myMap3 := map[int]string{
		1: "java",
		2: "c++",
		3: "python",
	}
	fmt.Println(myMap3)

}
