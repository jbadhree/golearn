package main

type Duration int64

const (
	ui         = 12345    // kind: integer
	uf         = 3.141592 // kind: floating-point
	ti int     = 12345    // type int
	tf float64 = 3.141592 // type float64
)

const (
	A1 = 3 << iota
	B1
	C1
)

func main() {
	println(ui, uf, ti, tf)
	var answer = 1 / 3
	println(answer)
	println(A1, B1, C1)

}
