package main

import ("fmt" 
		"math"
		"github.com/sapineni/go_crash_course/03_Packages/util"
		)

func main() 	{
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Pi)
	fmt.Println(math.Ceil(2.1))
	fmt.Println(util.ReverseString("hello"))
	fmt.Println(util.ConcatString("abc","xyz"))
}