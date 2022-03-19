package main

func main() {
	// Declate variable of type int with value of 10
	count := 10

	println("count:\tValue Of[", count, "]\t Addr Of[", &count, "]")

	// Pass the address of count
	increment(&count)
	println("count:\tValue Of[", count, "]\t Addr Of[", &count, "]")

}

// increment declares count as pointer variable
// value of this variable is always an address that points to type int
func increment(inc *int) {

	// Increment the value of count that the pointer points to
	*inc++
	println("inc:\tValue Of[", inc, "]\t Addr Of[", &inc, "]")

}
