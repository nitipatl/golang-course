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
