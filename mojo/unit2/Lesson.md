# Lesson Summary

## Declaring Variables

### Always Changing Variables:

- Non-Explicitly Typed Variable:
Use `var` keyword and assign it to a value:
`var num = 5`
*Note:* You must use the variable or you'll get a compilation error

- Explicitly Typed Variables
You specify the data type:
`var num2`

- Shorthand Declaration
Can use one of these:
<pre>
`name := 'mojo'`
`name, age := 'avi', 55`
`first, second string = mojo, avi // All must be of the same type`
`var name, age = 'mojo', 25 // allowed since data type was specified`
</pre>

<pre>
var (
    name = 'mojo'
    age = '55'
)
</pre>

*Note* You cannot use `:=` outside of a function. For that use `var`

### Never Changing Variables:

Use the `const` keyword:

`const name = 'mojo'`

## Removing Unused Variables:

Unused variables and packages will throw an error during compilation. Unused
variables declared outside a function will not throw an error though. To
declare an unused variable inside a function use: `_ = num`

## Dealing with Strings

Has escape characters like `\n and \t`

For raw strings, enclose it with:

<pre>
`This will print it
self as it
is. Similar to "'''" in
python`
</pre>

Strings can also be non english ones (like chinese or arabic)

## Type Conversion

### Discovering a Type

<pre>
fmt.Println(reflect.TypeOf(start)) // time.Time
fmt.Println(reflect.ValueOf(start).Kind()) // struct
</pre>

### Converting a variable's type

Consider:
<pre>
var age int
fmt.Print("Please enter your age: ") // Print doesn't add a new line
fmt.Scanf("%d", &age) // takes an interger *%d*, its similar to input()
// It will also assign the variable to age.
fmt.Println("You entered:", age)
</pre>
Above is not desirable because if a string is inputted the age will equal to
0. Below is a better approach:

<pre>
var input string
fmt.Print("Please enter your age: ")
fmt.Scanf("%s", &input) // input will have a string value

age, err := strconv.Atoi(input) //string converter package with function Atoi
(Ascii to interger) to convert the string to an interger.

if err != nil { // if there is an error
    fmt.Println(err)
} else { // if there is no error
    fmt.Println("Your age is:", age)
}

### Conversion Cheat Sheet

Use Parse to convert string to Bools, Floats or Intergers

<pre>
strconv.ParseBool('t')
strconv.ParseFloat("3.14", 64) // convert 3.14 to base64 float
strconv.ParseInt("-18", 10, 64) // convert to deci, 64bit
strconv.ParseUint("-18", 10,64) //gives error since its not Unassigned int
</pre>

To convert between different numerical types:

<pre>
num := 5
num2 := float32(num1)
num3 := float64(num2)
num4 := float32(num3)
num5 := int(num4)
</pre>

### String Formatting

You can use:
`s := name + 'a string' + strconv.Itoa(5) //convert interger to string`

A better solution is :

`s := fmt.Sprintf("%s, your number is %d", name, num)`

Here's a list of flags available:
<pre>
    d - decimal integer
    o - octal integer
    O - octal integer with 0o prefix
    b - binary integer
    x - hexadecimal integer lowercase
    X - hexadecimal integer uppercase
    f - decimal floating point, lowercase
    F - decimal floating point, uppercase
    e - scientific notation (mantissa/exponent), lowercase
    E - scientific notation (mantissa/exponent), uppercase
    g - the shortest representation of %e or %f
    G - the shortest representation of %E or %F
    c - a character represented by the corresponding Unicode code point
    q - a quoted character
    U - Unicode escape sequence
    t - the word true or false
    s - a string
    v - default format
    #v - Go-syntax representation of the value
    T - a Go-syntax representation of the type of the value
    p - pointer address
    % - a double %% prints a single %
</pre>

You can use a more General format without specifying the type:

<pre>
s := fmt.Sprintf("%v, your number is %v", name, num)
</pre>
