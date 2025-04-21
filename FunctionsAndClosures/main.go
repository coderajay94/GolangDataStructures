package main

import "fmt"

//functions and closures
func Print(x int) func() {
	return func() { fmt.Println("Printing from Print function", x) }
}

func main() {

	Print(50)()

	res := func(name string) {
		fmt.Println("Hello world")
	}

	func(name string) {
		fmt.Println("Hello world")
	}("ajay")

	//returns when function returns values, as anonymous function is called
	//without value can's assign anything
	out := func(name string) string {
		fmt.Println("Hello world")
		return "hello world" + name
	}("ajay")

	fmt.Print("out", out)

	res("Sds")
	res("SFG")

	c1 := counter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	var funcs []func() string

	for _, value := range []string{"ajay", "raghu", "yash"} {
		fmt.Println("adding the value", value)
		funcs = append(funcs, PrintMe(value))
	}

	for _, fun := range funcs {
		fmt.Println(fun())
	}

	//----------------------------
	//here for loop will take the last element for all
	values := []int{6, 1, 2, 3, 5}
	var prints []func()

	for _, value := range values {
		//if below line is missing it will take the for loop last value
		//You create a new variable temp scoped to that loop iteration,
		// so each closure captures a different copy of the value.
		temp := value
		prints = append(prints, func() {
			Printme(temp)
		})
	}

	for _, print := range prints {
		print()
	}

}

func Printme(x int) {
	fmt.Println("printing ", x)
}

func PrintMe(user string) func() string {
	return func() string {
		return "hello and welcome " + user
	}
}

func counter() func() int {
	count := 0
	return func() int {
		//if any value is updated inside and returned function
		//it creates a closure and share the same value for all the function calls
		// that means count increment values will be utilised if you call counter()() again
		//thats why in for loops only last value is used
		// to avoid this create local variables inside teh func and assign values but no updates
		count++
		return count
	}
}

//this will not update the count variable as it doesn't come inside the closure
func counter2() func() int {
	count := 0
	count++
	return func() int {
		ct := count
		return ct
	}
}
