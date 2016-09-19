
### Memory management

go is a low level language, which means we have to do a few things differently than in javascript or any high level language.   One of the most interesting and different things about Go compared to our experience with other languages is memory management.   Whenever you do something in Go, you allocate memory from the heap rather than from the stack.  When we make an array in javascript we write something like
```javascript
var a = [1,2,3,4,5]
```
javascript does three things for you when you make that array.   It allocates memory for the array, then uses the memory for the array when you call it.   Then it releases the memory when it isn’t needed anymore. 


In Go you need to be a little more verbose.   Here is what an array of numbers looks like in Go:
```go
 b := [5]int{1, 2, 3, 4, 5}
```
Because Go doesn’t always or have to manage memory for you, in an array you would need to specify the length.    This is to let the program knows how much memory to take from the heap.   Thankfully we don’t have to do this for most things in go.   Go has built in ways to manage memory.   This kind of memory management is just an option and something cool we haven’t seen yet.   Slices will have memory managed for them dynamically, and do the same thing as arrays.

Another cool thing about working with memory from the heap is that you can use pointers!   Here is what a pointer looks like
```go
a := foo

b := &a

*b = foo
```
When you set a variable you allocate a space of memory in the heap for that variable.  With a pointer you tell the program to look in that space.   Pointers can have they’re value resigned then along the way.   You can point to really anything in go.

### Typing in Go

Go is a staticly typed language, meaning that a string can't magically become an int, and an int can't magically become a string.

To declare a variable, one usually does something like this:
```go
var x float64 = 7.6
var y, z string = "abc", "123"
```
Now, x will always be a 64-bit floating point number, and y and z will always be a string. (NB: Declaring multiple vars with comma notation is amazing.) If one tries to change a type, one will get an error.
```go
var a int = 7
a = "harambe"
// cannot use "harambe" (type string)
// as type int in assignment

```
There are some shortcuts; Go will guess the variable type if not declared.
```go
var x = 7.6
var y = "abc"
```
Go will assume that 7 is a float, since there is a decimal between numbers, and that y is a string, since it is surrounded by double quotes.

This can be made even shorter with one of our favorite Go features: the `:=` operator.
```go
x := 7.6
y := "abc"
```

### Type Conversions
[Ref](https://golang.org/pkg/strconv)
#### Strings to Numbers
Go does allow us to change a few types with the `strconv` package. We used a couple of methods:
```go
xf, err := strconv.ParseFloat(x, 64)
```
`ParseFloat()` takes a string as its first argument and the width (in bits) of the float to be returned. It returns two values: a float of the specified width, and an error message. Note how we assign TWO variables to a single method. We actually skipped the error checking by using the nifty `_` variable name.
```go
xf, _ := strconv.ParseFloat(x, 64)
```
Using `_` in place of any variable name in Go causes it to be ignored, and no memory allocated for it. So if we wanted to ONLY check for errors, we could instead say:
```go
_, err := strconv.ParseFloat(x, 64)
```

#### And back again
After doing some math with our floats, we needed a string to print to the screen. For that, we used `FormatFloat()`, which is even more convoluted:
```go
ans := strconv.FormatFloat((xf * yf), 'f', -1, 64)
```
`FormatFloat()` returns a string, and takes in four arguments:
1. The float to be parsed (in this case a product)
* A byte representing the output format. We use the byte `'f'`, `102` in decimal or `01000110` in binary, to request a simple floating point number without any exponential notation.
* A precision value, in this case the number of digits after the decimal. Our value of `-1` lets Go decide how precise it needs to be in order to be accurate.
* The bit width of the original float, in this case 64.

There are some easier methods with integers (`atoi` may be familiar to C fans) and strings.

### Reading Files in Go

To serve up our `index.html` page, we had to do a tiny bit of os-level stuff. We used two additional native packages, `os` and `io/ioutil`.

The `net/http` server package of Go writes a slice of bytes to the response body, so we couldn't natively just say, "Hey go, give them this file on the server!". To serve up the `index.html`, we used the following:
```go
func IndexHandler(w http.ResponseWriter, r *http.Request) {
  index, err1 := os.Open("index.html")

  if err1 != nil {
    panic(err1)
  }

  data, err2 := ioutil.ReadAll(index)

  if err2 != nil {
    panic(err2)
  }

  w.Write([]byte(data))
}
```
`os.Open()` returns a file pointer and an error message, and takes in a string representing a relative path file location.

`ioutil.ReadAll()` returns a byte slice (`[]byte`) and an error message, and takes in a file pointer.

`http.responseWriter.Write()` appends the given `[]byte` to the response body.

In this example, we don't really "catch" errors, but instead use `panic()` to immediately halt operation if an error is returned. It is the best-named function of all time.

### Using Mux
Mux is Go package that implements a request router and dispatcher. The method, `Mux.NewRouter()`, matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions. This can be seen in the following code:

```Go
func main() {
   r := mux.NewRouter()
   fmt.Println("Listening on port 8000\n")
   // Routes consist of a path and a handler function.
   r.HandleFunc("/", IndexHandler)
   r.HandleFunc("/add", AddHandler)
   r.HandleFunc("/sub", SubHandler)
   r.HandleFunc("/mult", MultHandler)
   r.HandleFunc("/div", DivHandler)
   //Bind to a port and pass our router in
   log.Fatal(http.ListenAndServe(":8000", r))
}
```

In our code, we assign `NewRouter()` to r and use a handler to match the url, as well as retrieve a `[]byte` from a function.


