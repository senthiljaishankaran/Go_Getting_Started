// Declares that the current Go source file is part of an executable program
// The main package is special because it signifies the entry point for your program's execution
package main

// fmt is the package contains functions for formatted input/output operations.
import (
	"fmt"
)

func main() {
	fmt.Println("First Go Program")

	fmt.Println("Variable Declaration and type definition")
	var first_name string = "Senthil" // type declared variables
	var last_name string = "Kumar"    // type declared variables
	var age int64 = 30                // type declared variables
	var isLoggedIn bool = true        // type declared variables
	age = 32                          // mutable

	fmt.Println("User Name is", first_name+last_name)
	fmt.Println("Is he logged in:", isLoggedIn)
	fmt.Println("His age is:", age)

	// Inferred type short variable declaration
	newName := "Suresh Kumar"
	newAge := 31
	isAlive := true
	fmt.Println("New Name is:", newName)
	fmt.Println("His Age is:", newAge)
	fmt.Println("isAlive:", isAlive)

	fmt.Println("Constants")
	const pie float64 = 3.14
	fmt.Println("Value of pie is:", pie)

}
