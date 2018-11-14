package main

import "fmt"
var sal = 10.5
func main(){
	name := "Sapi"
	age:= 10
	name = name + " Rao"
	fmt.Println("Name is ", name)
	fmt.Println("age is ",
	age)
	first_name, lastName := "john", "doe"
	fmt.Printf("%T %T %T\n",name,age,sal)
	fmt.Printf("first name is %s and last name is %s",first_name,lastName)
	test()
}
func test(){
	sal += 10
	fmt.Println("sal is ",sal)
}