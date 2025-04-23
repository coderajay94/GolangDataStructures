// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	values := []int{12, 23, 11}
	var funcs []func()
	for _, value := range values {
		//with version 1.22 every value coming from 
		//for loop is new variable
		//avoiding issue of returning same value again and again
		funcs = append(funcs, func() {
			fmt.Println("hellow from closure with value:", value)
		})
	}

	for _, fun := range funcs {
		fun()
	}
}
