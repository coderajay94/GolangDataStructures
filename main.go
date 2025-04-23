// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {

	name := "ajay kumar"
	var reverse string
	for _, ch := range name {
		reverse = string(ch) + reverse
	}

	fmt.Println("reverse is printed", reverse)
}
