**Return multiple values**  

```go
func f() (int, int) {
     return 5, 6
}
func main() {
     x, y := f()
}
```

Multiple values are often used to return an error value along with the result (x, err := f()), or a boolean to indicate success (x, ok := f()).   

**Variadic Functions**  
```go
func add(args ...int) int {
     total := 0
     for _, v := range args {
           total += v
}
     return total
}
func main() {
     fmt.Println(add(1,2,3))
}
```
Ex. fmt.Println  

Convert to commas  

```go
func main() {
     xs := []int{1,2,3}
     fmt.Println(add(xs...))
}
```

**Closure**  

```go
func main() {
     add := func(x, y int) int {
            return x + y }
     fmt.Println(add(1,1))
}
```

```go
func makeEvenGenerator() func() uint {
     i := uint(0)
     return func() (ret uint) {
            ret = i
            i += 2
            return 
    }
}
func main() {
     nextEven := makeEvenGenerator()
     fmt.Println(nextEven()) // 0
     fmt.Println(nextEven()) // 2
     fmt.Println(nextEven()) // 4
}
```

**Defer, Panic & Recover**  
defer = schedules a function call to be run after the function completes.  
```go
f, _ := os.Open(filename)
defer f.Close()
```

Panic & Recover  
panic => then combo with defer 