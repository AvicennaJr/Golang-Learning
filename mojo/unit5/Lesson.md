## Defining a Function

<pre>
package main

import (
    "fmt"
    "time"
    )

func displayDate() {
    fmt.Println(time.Now().Date())
}

func main() {
    displayDate()
}

</pre>

### With parameters:

<pre>
func displayDate(format string) {
    fmt.Println(time.Now().Format(format))
}
</pre>

And to call it:

`displayDate("Mon 2006-01-02 15:04:05")`

How go formats time //very unnecessary tbh

<pre>


Reference Value                 Description               Example Usage

Mon                           Day of the week                Mon, Monday

Jan                                 Month                   Jan, January, 01

2                                   Day                         2, 02

15                                  Hour                        15, 3

04                                  Minute                      04, 4

05                                  Second                      05, 5

2006                                Year                        2006

MST                                 Time zone                   MST

</pre>

### With multiple parameters

<pre>

func displayDate(format string, prefix string) {
    fmt.Println(prexfix, time.Now().Format(format))
}

displayDate("Mon 2006", "Current Day and year")

</pre>

### Passing Arguments by value

Lets say you want to pass in arguments to a function without changing the
values of the arguments, do it like this:

<pre>

func swap(a, b int) {
    a, b = b, a
}

func main() {
    x := 5
    y := 6

    swap(x, y) 

    fmt.Println(x, y) // remains 5, 6
}

</pre>

### Passing Arguments by pointer

Now if you want to change the values of x and y then:

<pre>

func swap(a, b `*int`) {
    *a, *b = *b, *a
}


swap(&x, &y)

fmt.Println(x, y) // 6 and 5

</pre>

### Returning Values

<pre>

func addNum(num1, num2 int) int { // the return value is int
    return num1 + num2
}

s := addNum(5,6)
fmt.Println(s) //11

</pre>

You can also return multiple values:

<pre>

func countOddEven(s string) (int, int) { // the returning values are both int
    odds, evens := 0,0

    for _, c := range s {
        if int(c) % 2 == 0 {
            evens++
        } else {
            odd++
        }
    }

    return odds, evens
}

o, e := countOddEven('12345')

fmt.Println(o, e) // 3,2

// if you just want evens:

_ , e := countOddEven('12345')

// if you just want odds:

o, _ := countOddEven('12345')
</pre>

### Naming Return Vaules

<pre>

func addNum(num1 int, num2 int) (sum int) {
    sum = num1 + num2

    return // will return sum
}

</pre>

### Variadic Functions

<pre>

func addNums(nums ... int) int {
    total := 0

    for _, n := range nums {
        total += n

    }
    return total
}

fmt.Println(addNums(1,2,3,4,5) // 15
fmt.Println(addNums(1,2,3) // 6

// Also 
func addNums(total int, nums ... int) int { // variadic must always be last

    fmt.Printf(%T", nums)
    for _, n := range nums {
        total += n

    }
    return total
}

</pre>

## Anonymous Functions

Function without a name

### Declaration

<pre>

func main() {
    var i func() int
}

//Declared i to be a function that returns int

i = func() int {
    return 5
}

fmt.Println(i())

</pre>

### Closures

This is an advanced concept. A closure is a function value that references
variables from outside its body.

<pre>

func fib() func() int { // returns a function that returns an int
    f1 := 0
    f2 := 1

    return func() int { // it is basically returning an anonymous function
        f1, f2 = f2, (f1 + f2)

        return f1
    }

}

</pre>

What makes the above a closure is that the returning functions can access f1
and f2 which are declared outside its body.

You can assign the fib() function to a variable:

<pre>

func main() {
    gen := fib() // gen is basically the returning function from fib()
}

// call gen()

fmt.Println(gen()) // 1 - first fibonacci number
fmt.Println(gen()) // 1
fmt.Println(gen()) // 2

</pre>

It is useful because f1 and f2 declared outside hold the values of the last two
numbers generated. To ask the next ten numbers of fibonacci:

<pre>

func main() {
    gen := fib()
    for i := 0; i < 10, i++ {
        fmt.Println(gen())
    }
    }

</pre>


### Implementing Filter Functions using Closures

I was wondering where closures can be implemented, turns out that they make it easy to
implement filter functions, mapping functions etc. Heres an example:

<pre>

// create a filter function that will return values of a number that meet a
// certain requirement:

func filter(arr [] int, cond func(int) bool) []int { // the first parameter is
// an array of numbers, the second parameter is a closure function that takes
// in int and returns a bool

    result := []int{} // create a new array that we'll append values that meet
    // the required conditions

    for _, v := range arr {
        if cond(v) { // if the number satisfies the condition
            result = append(result, v) // append the number to result
        }
    }

    return result // return the values we want
}

</pre>

Let's use the function we made to filter out only even numbers:

<pre>

func main() {
    a := []int{1,2,3,4,5}

    evens := filter(a,
        func(val int) bool {
            return val%2 == 0 // returns true if divisible
        })

    fmt.Println(evens) // [2,4]
}

</pre>

Lets say we want numbers that are less than 5:

<pre>

less5 := filter(a,
        func(val int) bool {
            return val < 5 // true is a value is less than 5
        })

fmt.Println(less5) // [1,2,3,4]

</pre>

Its complicated but easy.

> And we're done with functions
