**Variables**  
```var x string = "Hello World"```  
no type was speci- fied. The type is not necessary because the Go compiler is able to infer the type based on the literal value you assign the variable.  
```var x = "Hello World"```  
```x := "Hello World"```   

**How to Name a Variable**  
1. Pick names which clearly describe the variable's purpose.  
```go
x := "Max"
fmt.Println("My dog's name is", x)
```
x is not a very good name for a variable  
```go
name := "Max"
fmt.Println("My dog's name is", name)
```
or
```go
dogsName := "Max"
fmt.Println("My dog's name is", dogsName)
```

2. Scope  
```go
package main
import "fmt"
var x string = "Hello World"
func main() {
    fmt.Println(x)
}
```  
This means that other functions can access this variable of `x`

```go
func main() {
    var x string = "Hello World"
    fmt.Println(x)
}
func f() {
    fmt.Println(x)
}
```
The range of places where you are allowed to use x is called the scope of the variable.  

3. Constants  
var -> const  
```go
const x string = "Hello World"
x += "OK"
```
Output : `cannot assign to x`  
Constants are a good way to reuse common values in a program without writing them out each time.  

4. Defining Multiple Variables  
```go
var ( 
a = 5
b = 10
c = 15 )
```
`var` and `const` can be used  


**Problems**  
1. What are two ways to create a new variable?   
```go
var x string = "OK"
var x = "OK"
x := "OK"
```
2. What is the value of x after running: `x := 5; x += 1`?  
x = 6  
3. What is scope and how do you determine the scope of a variable in Go?  
```go
main() {
	// this scope can be use x variable
	x := "OK"
	f()
}
f() {
	fmt.Println(x) // <~ can't access x variable because out of scope
} 
```
4. What is the difference between var and const?  
variable = value can be set everytime after intitial  
const = value can't change again after intitial ex. PI  

5. Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)  

6. Write another program that converts from feet into meters. (1 ft = 0.3048 m)