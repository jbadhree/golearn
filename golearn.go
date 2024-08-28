package main

import (
	"fmt"
	"os"

	algods "github.com/jbadhree/golearn/algods"
	basics "github.com/jbadhree/golearn/basics"
)

func main() {
	// Get first cli argument and store it in a string
	cli := os.Args[1]
	fmt.Println("CLI argument is", cli)
	switch cli {
	case "arrays":
		basics.ExecuteArrays()
	case "runlength":
		algods.ExecRunLengthEncoding()
	case "mergesort":
		arr := []int{5, 1, 7, 3, 8, 9}
		fmt.Println(algods.MergeSort(arr))
	}

}
