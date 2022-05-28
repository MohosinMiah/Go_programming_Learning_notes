package main

import (
	"fmt"
)
func main() {
	var n [10] int /* n is an array of 10 integers */
	var i,j,k int

	var friends = [] string{"MOHOSIN","SHARGO"}
 
	/* initialize elements of array n to 0 */         
	for i = 0; i < 10; i++ {
	   n[i] = i + 100 /* set element at location i to i + 100 */
	}
	
	/* output each array element's value */
	for j = 0; j < 10; j++ {
	   fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}

	/* Out Put Array */

	for k = 0; k < len(friends) ; k++ {
		fmt.Println(friends[k])
	}
 }