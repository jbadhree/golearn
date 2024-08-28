package basics

func ExecPassByValue() {
	// Declate variable of type int with value of 10
	count := 10

	println("count:\tValue Of[", count, "]\t Addr Of[", &count, "]")

	// Pass the value of count
	Increment(&count)
	println("count:\tValue Of[", count, "]\t Addr Of[", &count, "]")

}

// // increment the value passed
// func increment(inc int) {
// 	// Increment the value of inc
// 	inc++
// 	println("inc:\tValue Of[", inc, "]\t Addr Of[", &inc, "]")

// }
