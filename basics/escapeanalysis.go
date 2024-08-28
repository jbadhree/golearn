package basics

// Represents a user in the system
type user struct {
	name  string
	email string
}

func ExecEscapeAnalysis() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)

}

// factory function that creates a user and returns it to the user
// value semantics
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "abc@sbc.com",
	}

	println("V1", &u)

	return u
}

// factory function that creates a user and returns it to the user
// pointer semantics
func createUserV2() *user {
	// Clever Code
	// This will work but return statement will not have &
	// And it will not tell just by looking at return if the value is in heap or stack
	/*
		u := &user{
			name:  "Bill",
			email: "abc@sbc.com",
		}
	*/

	// Good Code
	u := user{
		name:  "Bill",
		email: "abc@sbc.com",
	}

	println("V2", &u)

	// Clever Code's return statement
	// return u

	// Good Code return statement increased readability
	return &u
}

/*
// Bad V2 Function
func createUserV2() user {

	u := &user{
		name:  "Bill",
		email: "abc@sbc.com",
	}

	println("V2", &u)

	return *u
}
*/
