# Consuming WEB APIs

## Fetching Data

You can use the Get() function from the net/http package. Get returns a
response struct and an error. You can also use the ReadAll() function from the
io/iotil package to read the response from the server. ReadAll() returns a
slice of bytes from the server and an error.

Example Code can be found on main.go.
