package basics

import "fmt"

type sliceUser struct {
	likes int
}

func ExecSlices() {

	/*
		fruits := make([]string, 5)
		fruits[0] = "Apple"
		fruits[1] = "Orange"
		fruits[2] = "Banana"
		fruits[3] = "Grape"
		fruits[4] = "Plum"
	*/

	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	//var data []string

	data := make([]string, 1e5, 1e5)

	lastCap := cap(data)
	// fmt.Printf("Last Cap before Loop = %d\n",lastCap)
	// Append 100K strings to the slice
	//for record := 1; record <= 1e5; record++ {

	for record := 0; record < 1e5; record++ {

		// Use built in function to append to slice
		value := fmt.Sprintf("Rec: %d", record)
		//data = append(data,value)

		data[record] = value

		if lastCap != cap(data) {
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100
			lastCap = cap(data)
			fmt.Printf("Addr[%p] Index[%d] Cap[%d  - %2.f %% ]\n", &data, record, cap(data), capChg)
		}

	}

	//inspectSlice(fruits)// operating on its own copy of slice

	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1)

	// this is going to take elem from [2] to [4-1]=[3]
	slice2 := slice1[2:4:4]
	inspectSlice(slice2)

	//slice2[0] = "CHANGED"
	slice2 = append(slice2, "CHANGED")

	inspectSlice(slice1)
	inspectSlice(slice2)

	// copy function
	// copy slice1 to a new slice3
	slice3 := make([]string, len(slice1))
	copy(slice3, slice1)

	inspectSlice(slice3)

	// code to demostrate why append creating its own backing array could be dangerous

	// declare a slice of 3 users
	users := make([]sliceUser, 3)

	// WRONG share address user at index 1
	//shareUser := &users[1]

	// RIGHT
	users[1].likes++

	// WRONG add a like to shareUser
	//shareUser.likes++

	// Display the number of likes for all users.
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", users[i], users[i].likes)
	}

	// Add a new user - this will create a copy of the original slice and operate on that
	users = append(users, sliceUser{})

	// WRONG add another like to shareUser.. but this it still pointing to original backing array
	//shareUser.likes++

	// RIGHT Add another like
	users[1].likes++

	// Display the number of likes for all users.
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", users[i], users[i].likes)
	}

}

func inspectSlice(slice []string) {
	fmt.Printf("Length [%d] Capacity [%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s %s\n", i, &slice[i], slice[i], s)
	}
}
