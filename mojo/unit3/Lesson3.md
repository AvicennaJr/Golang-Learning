# Making Descisions

## IF/ELSE Statements

### Comparison Operators

<pre>
Operator  Description
    ==      Equal to
    !=      Not equal to
    <       Less than
    <=      Less than or equal to
    >       Greater than
    >=      Greater than or equal to
</pre>

Combining Operators:
<pre>
    &&      x && y
    ||      x || y
    !       !x
</pre>

### if/else Statements

<pre>
num := 6

if num % 2 == 1 {
    fmt.Println("Number is odd")
} else {
    fmt.Println("Number is even")
}
</pre>

### Short Circuiting

Similar to Python

<pre>

func raining() {
    fmt.Println("Its raining Now")
    return true
}

func snowing() {
    fmt.Println("Its snowing now")
    return true
}

if ranining() || snowing() {
    fmt.Println("Stay inside")
} // will print raining() and skip snowing

if !raining() || snowing() {
    fmt.Println("Lets ski")
} // since raining() is false,it'll check snowing()

if raining() && snowing() {
    fmt.Println("It'll be coooooooooooold")
} // both have to be true (&&)

</pre>

### Using Variables in if/else scope

You can define a function that returns two variables like this:

<pre>

func doSomething() (int, bool) { // return an integer and bool
    return 5, false
}
</pre>

Then you can assign the returning values to two variables like this:

<pre>

v, err := doSomething()

if err {
    // if error is not false handle the error
} else { // if no errors then proceed
    fmt.Println(v)
}
</pre>

Notice that `v` and `err` can be accessed outside the if/else block. If you
want to limit their scope to just the if/else block and not be able to be
accessed outside then use this:

<pre>

if v, err := doSomething(); !err { // if there's no errors
    fmt.Println(v)
} else {
    // handle the error
} 
</pre>

This way the variables `v` and `err` can't be used anywhere else.

### Ternary Operators

This is the Go way to write ternary operators:

`By using if/else statements *ba dum tss*`

To make a 'custom' ternary operator do this:
<pre>
func checkParity(num, int) string { //accept int as input and return string
    if num % 2 == 0 {
        return 'even'
    } else {
        return 'odd'
    }
}

num := 5

parity := checkParity(num)

fmt.Println(parity) // output is odd

</pre>

## Switch Statement

Use this if you have multiple conditional statements (think if/elseif/else):

<pre>
num := 5

day := ""

switch num {
    case 1: // if num == 1
        day = "Mon"
    case 2:
        day = "Tue"
    case 3:
        day = "Wed|
    case 4:
        day = "Thur"
    case 5:
        day = "Fri"
    case 6:
        day = "Sat"
    case 7:
        day = "Sun"
    default: // if num is not in these options
        day = "----error----"
}

fmt.Println(day) // output is Fri

</pre>

 > Notice it doesn't need a break statement like Javascript

You can have multiple statements to execute for each case, like:

<pre>

case 1:
    day = 'Mon'
    fmt.Println("I hate Mondays")
</pre>

### Switch Fallthroughs

To continue matching cases after a particular match, use `fallthrough` keyword:

<pre>

grade := 'C'

switch grade{
    case 'A':
        fallthrough
    case 'B':
        fallthrough
    case 'C':
        fallthrough
    case 'D':
        fmt.Println("Passed")
    case 'F':
        fmt.Println("Failed")
    default:
        fmt.Println("Absent")
}

</pre>

### Matching Multiple Cases

The previous snippet can be re-written as this:

<pre>

switch grade {
    case "A", "B", "C", "D":
        fmt.Println("Pass")
    case "F":
        fmt.Println("Fail")
    default:
        fmt.Println("Absent")
}
</pre>

### Switching without Conditions

You can write switch statements without using any conditions as follows:

<pre>

score := 79

grade := ""

switch { // note that if score isn't an int it'll throw an error:
    case score < 50: grade = "F"
    case score < 60: grade = "D"
    case score < 70: grade = "C"
    case score < 80: grade = "B"
    case score >= 80: grade = "A"
    default: fmt.Prinln("Undefined Score")
}

fmt.Println(grade)

</pre>

*That's all about Conditions*
