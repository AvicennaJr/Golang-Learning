package main

import (
	"fmt"
	"time"
)

func displayDate() {
	fmt.Println(time.Now().Date())
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func countOddEven(s string) (int, int) { // the returning values are both int
	odds, evens := 0, 0

	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}

	return odds, evens
}

func fib() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f1, f2 = f2, f1+f2

		return f1
	}
}

func filter(arr []int, cond func(int) bool) []int {

	results := []int{}

	for _, v := range arr {
		if cond(v) {
			results = append(results, v)
		}
	}

	return results
}

func main() {
	displayDate()
	x := 5
	y := 6

	swap(&x, &y)

	fmt.Println(x, y)
	_, e := countOddEven("12345")
	fmt.Println(e)

	gen := fib()

	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}

	// filter out numbers between 10 and 20

	a := []int{1, 12, 3, 25, 16, 9}

	var_filter := filter(a,
		func(val int) bool {
			return val > 10 && val < 20
		})

	fmt.Println(var_filter)
}
