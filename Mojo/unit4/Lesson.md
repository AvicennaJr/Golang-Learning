# Loops

For statement looks like this:

<pre>
for (init; condition; post) {
}
</pre>

- init: Statement executed before the first iteration starts
- condition: Expression evaluated before the iteration starts to determine if
  the iteration should continue
- post: Statement to be evaluated at the end of each iteration

Example:
<pre>
package main
import 'fmt'

func main() {
    for i:=0; i<5; i++ { // notice there's no brackets
        fmt.Println(i)
    }
}
</pre>

## Incremental Operators

`x++` - this will add 1. Equivalent to x = x + 1
`x--` - this will subtract 1. Equivalent to x = x - 1

## Fibonacci Example

You can write a for loop with just the `codition` (ommiting the `init` and
`post`)

<pre>
max := 100
a, b := 0, 1
for b <= max {
    println(b)
    a, b = b, a+b
}
</pre>

*Notice that the code above is technically a while loop*

## Infinite Loops

Go does not have `while` loops. An infinite loop in Go is written as:
<pre>
for {
}
</pre>

Example:
<pre>
package main

import (
    "fmt"
    "strings"
    )

func main() {
    for {
        fmt.Println("Enter QUIT to exit:")
        var input string
        fmt.Println("Please enter a word:")
        fmt.Scanln(&input)
        if strings.ToUpper(input) == "QUIT" {
            break
        }
    }
}
</pre>

Another Example:
<pre>
for n := 1; n<10; n++ {
    if n%2 == 0 {
        continue // skip the remaining portion and start a new iteration
    }
    fmt.Println(n)
}

</pre>

Alternative to Fibonacci:
<pre>
max := 100

for a, b := 0; 1 < b <= max; a, b = b, a+b {
    println(b)
}
</pre>


## Iterating over a Range of Values

Used when iterating through a collection of items.

### Iterating through `arrays/slices`

<pre>
var OS [3]string // you'll learn about arrays soon
OS[0] = 'iOS'
OS[1] = 'Android'
OS[2] = 'Windows'

for i, v := range OS {
    fmt.Println(i, v)
}

/* 
Out put is:
0 iOS
1 Android
2 Windows
*/
</pre>

Notice that it returns both the index and its value. If you don't want the
index then use the blank identifier:

<pre>

for `_`, v := range OS {
    fmt.Println(v)
}
</pre>

To get just the indexes:
<pre>
for i := range OS {
    fmt.Println(i)
}
</pre>

### Iterating through a `string`

<pre>
for pos, char := range "Hello World" {
    fmt.Println(pos, char)
}

// This will print out unicode of each character. If you want to get the actual
// character use:

for pos, char := range "Hello World" {
    fmt.Printf("%d %c\n", pos, char)
}
</pre>

## Using Labels with `for` loop

consider the following:

<pre>

for i := 1; i <= 5; i++ {
    for j := 1; j <=5; j++ {
        fmt.Printf("%d * %d = %d\n", i, j, i*j)
    }
    fmt.Println("--------------")
}

// this is the output


1 * 1 = 1
1 * 2 = 2
1 * 3 = 3
1 * 4 = 4
1 * 5 = 5
-----------
2 * 1 = 2
2 * 2 = 4
2 * 3 = 6
2 * 4 = 8
2 * 5 = 10
-----------
3 * 1 = 3
3 * 2 = 6
3 * 3 = 9
3 * 4 = 12
3 * 5 = 15
-----------
4 * 1 = 4
4 * 2 = 8
4 * 3 = 12
4 * 4 = 16
4 * 5 = 20
-----------
5 * 1 = 5
5 * 2 = 10
5 * 3 = 15
5 * 4 = 20
5 * 5 = 25
-----------
</pre>

Suppose you only want to print the multiplication table of only 1 and 2:

<pre>
for i := 1; i<=5; i++ {
    for j := 1; j <=5; j++ {
        if i == 3 {
            break
        }
        fmt.Printf("%d * %d = %d", i, j, i*j)
    }
    fmt.Println('-------------')
}

// thiswill be the output


1 * 1 = 1
1 * 2 = 2
1 * 3 = 3
1 * 4 = 4
1 * 5 = 5
-----------
2 * 1 = 2
2 * 2 = 4
2 * 3 = 6
2 * 4 = 8
2 * 5 = 10
-----------
-----------
4 * 1 = 4
4 * 2 = 8
4 * 3 = 12
4 * 4 = 16
4 * 5 = 20
-----------
5 * 1 = 5
5 * 2 = 10
5 * 3 = 15
5 * 4 = 20
5 * 5 = 25
-----------
</pre>


What you actually intended was to exit the two for loops totally — both the inner
as well as the outer loops. But the break statement only breaks out from the
innermost loop that it’s in. To fix this, you need to use a label for your outer loop,
and then specify where the break statement breaks out to:

<pre>


Outerloop:
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            if i == 3 {
                break Outerloop
            }
            fmt.Printf("%d * %d = %d\n", i, j, i*j)
        }
    fmt.Println("-----------")
    }
}

// This is the output

1 * 1 = 1
1 * 2 = 2
1 * 3 = 3
1 * 4 = 4
1 * 5 = 5
-----------
2 * 1 = 2
2 * 2 = 4
2 * 3 = 6
2 * 4 = 8
2 * 5 = 10
-----------
</pre>

You could also use the `continue` statement:

<pre>


Outerloop:
    for i := 1; i <= 5; i++ {
        for j := 1; j <= 5; j++ {
            if i == 3 {
                continue Outerloop
            }
            fmt.Printf("%d * %d = %d\n", i, j, i*j)

            fmt.Println("-----------")
            }
        }
    }

// output is 


1 * 1 = 1
1 * 2 = 2
1 * 3 = 3
1 * 4 = 4
1 * 5 = 5
-----------
2 * 1 = 2
2 * 2 = 4
2 * 3 = 6
2 * 4 = 8
2 * 5 = 10
-----------
4 * 1 = 4
4 * 2 = 8
4 * 3 = 12
4 * 4 = 16
4 * 5 = 20
-----------
5 * 1 = 5
5 * 2 = 10
5 * 3 = 15
5 * 4 = 20
5 * 5 = 25

</pre>

> That's all about Loops in Golang
