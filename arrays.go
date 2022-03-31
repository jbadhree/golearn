package main

import "fmt"

func main() {

	var fruits [5]string
	fruits[0] = "Apple" // Assignment is a copy operation
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	for i, fruit := range fruits {
		fmt.Println(i, fruit)

	}

	fmt.Print("Pointer Semantics\n")
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Before[%s] : ", friends[1])

	// for.. range using Pointer Semantics
	// when we say range friends, it points to the variable by default since friends is an array
	// Here the range method points to the friends array and operates on it directly
	for i := range friends {
		friends[1] = "Jack"
		if i == 1 {
			fmt.Printf("After[%s]\n", friends[1])
		}
	}

	fmt.Print("Value Semantics\n")

	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Before[%s] : ", friends[1])

	// for.. range using Value Semantics
	// Here the range copies the friends array and iterates over it
	// And hence when friends[1] is updated, it doesnt actually update the copy which is being iterated over
	for i, v := range friends {
		friends[1] = "Jack"
		if i == 1 {
			fmt.Printf("After[%s]\n", v)
		}
	}

	fmt.Print("Mixing Semantics - DONT DO THIS\n")

	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Before[%s] : ", friends[1])

	// for.. range Mixing Semantics
	// Here i, v is used, which means range method uses value semantics
	// But since &friends is used. it actually does not copy, it operates over the original friends array
	// Hence v actually points to the friends[i] value
	for i, v := range &friends {
		friends[1] = "Jack"
		if i == 1 {
			fmt.Printf("After[%s]\n", v)
		}
	}

	// See how memory is allocated
	// See how the memory is contiguous
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
	}
}
