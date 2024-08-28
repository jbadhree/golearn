package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Print("Error")
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	for _, line := range lines {
		fmt.Println(line)
	}
}
