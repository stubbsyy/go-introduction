package main

import "fmt"

func main() {
	var number = 10
	var numberTwo int
	numberThree := 15
	moveForward := true
	var moveBack = false
	var helloTHere string = "hi"
	goodBye := "See you tomorrow"
	numberTwo = 12
	const thisCantChange = 25
	var thisCanChange = 25
	fmt.Println(number)
	fmt.Println(numberTwo)
	fmt.Println(numberThree)
	fmt.Println(moveForward)
	fmt.Println(moveBack)
	fmt.Println(helloTHere)
	fmt.Println(goodBye)
	//thisCantChange *= 2
	thisCanChange *= 2
	fmt.Println(thisCanChange)
	fmt.Println(thisCanChange*2, "The end")
}