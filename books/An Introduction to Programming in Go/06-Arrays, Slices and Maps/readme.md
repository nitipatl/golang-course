**Range**
```go
var total float64 = 0
for _, value := range x {
     total += value
}
fmt.Println(total / float64(len(x)))
```

A single _ (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator variable)