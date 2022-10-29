# Packages

Go organizes code into units call packages. A package is made up of a collection of
files. The main package is a special package that contains the main() function,
which makes the main package an executable program. The main() function
serves as the entry point to your application.

*All files in a package must be in the same directory, and all package names must
all be in lowercase.*

*Also to run a program that has multiple packages (eg main.go, point.go) use go
run * .go*

## Creating Shareable Packages

You can add the code into your go path (/home/mojo/go) inside a folder called
src. So the structure is /home/mojo/go/src/geometry/point.go. Install the package
with `go install` (while being inside the geometry folder) Then you can import the
geometry package like "fmt" and use it like `pt1 := geometry.Point{X:1, Y:2}`.

*Note that point.go will be of package geometry*

## Organizing Packages into directories

Lets say you wanted to divide geometry into sub packages like coordinates etc.
Then create a new folder under geometry call it coordinate and place the
point.go file under it. Now the package name of point.go will be `package
coordinate`. Go to the coordinate folder and install it with `go install`. You
can then use it in other codes like `import "geometry/coordinates"`

## Using third partry packages

Go doesn't have a central repository like Pypi or NPM. Instead you fetch third
party packages through a hostname and path like Github. You use the command "go
get /path/to/package".
