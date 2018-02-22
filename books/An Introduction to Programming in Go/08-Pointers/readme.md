**Problems**  
1. How do you get the memory address of a variable?  
&variable
2. How do you assign a value to a pointer? 
*variable = value  
3. How do you create a new pointer?
a := new(int)  
4. What is the value of x after running this program: 
```go
func square(x *float64) {
    *x = *x * *x
}
func main() {
    x := 1.5
    square(&x) 
}
```
1.5*1.5 = 2.25  
