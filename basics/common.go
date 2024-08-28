package basics

func Increment(inc *int) {

	// Increment the value of count that the pointer points to
	*inc++
	println("inc:\tValue Of[", inc, "]\t Addr Of[", &inc, "]")

}
