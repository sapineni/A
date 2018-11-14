package main

import "fmt"

func main() {
	var stringArr [2]string
	stringArr[0] = "Hello"
	stringArr[1] = "World"
	stringFruit := [2]string {"Apple", "Banana"}
	fmt.Println(stringArr)
	fmt.Println(stringFruit)

	furitSlice := []string {"Pink","Green","Red"}
	
	fmt.Println(furitSlice)

	fmt.Println(furitSlice[1])
	fmt.Println(furitSlice[1:3])
}