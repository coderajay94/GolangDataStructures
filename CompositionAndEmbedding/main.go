package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	TempAddress TempAddress `json:tempAddress` //composition
	PermAddress             //embedding - direclty have fields in to parent, if found same will igore to avoid conflict
}

type TempAddress struct {
	City string `json:"city"`
}

type PermAddress struct {
	City string `json:"city"`
}

func main() {

	fmt.Println("Hello world")

	emp1 := Employee{
		FirstName: "akay",
		LastName:  "kumar",
		TempAddress: TempAddress{
			City: "Bangalore",
		},
		PermAddress: PermAddress{
			City: "Abohar",
		},
	}

	fmt.Println(emp1)

	resp, _ := json.Marshal(emp1)

	fmt.Println(string(resp))

	values := []int{1, 2, 3, 5}
	var prints []func()

	for index, value := range values {
		prints = append(prints, func() {
			fmt.Println("adding value", value, index)
			Printme(value)
		})
	}

	for _, print := range prints {
		print()
	}

}

func Printme(x int) {
	fmt.Println("printing ", x)
}
