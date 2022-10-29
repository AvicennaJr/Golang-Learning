# Grouping Packages into Modules

A module is a directory of packages with a file named go.mod at its root. The
go.mod file defines the modules path and dependency requirements.

## Creating and Publishing a Module

<pre>


$HOME
    |__stringmod
        |__strings
            |__strings.go
        |__quotes
            |__quotes.go

</pre>

Just do `go mod init *module name*` and it will create a module.

To publish the module on Github, just initialize the module as above and upload
it to github. It will be readily available for everyone to use.
