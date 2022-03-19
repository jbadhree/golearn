// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// Named types
type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	fmt.Println("Hello")

	// Literal types
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141692,
	}

	var a bill
	var b alice
	a = e2
	fmt.Println(a, b, e2)
}
