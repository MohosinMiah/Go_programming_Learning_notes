package main

import "fmt"

func main() {

	nameOne := "Bangladesh"
	nameTwo := "India"

	x, y := swap(nameOne, nameTwo)

	fmt.Println("Swap Name ")
	fmt.Println(x)
	fmt.Println(y)

}

func swap(nameOne, nameTwo string) (string, string) {
	return nameTwo, nameOne
}