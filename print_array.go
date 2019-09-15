package main

import "fmt"

// PrintArray 配列をPrintする
func PrintArray() {
	a := [4]int{1, 2, 3, 4}
	b := a[3]

	fmt.Println(b)

	fmt.Println(a)

	q, r := div(0, 0)
	fmt.Printf("%d, %d", q, r)

	fmt.Println("")
	q1, r1 := doSomething()
	fmt.Printf("%d, %d", q1, r1)
}

func div(x, y int) (int, int) {
	return 0, 0
}

func doSomething() (x, y int) {
	y = 5
	return
}
