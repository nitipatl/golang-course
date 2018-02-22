**Range**
```go
var total float64 = 0
for _, value := range x {
     total += value
}
fmt.Println(total / float64(len(x)))
```

A single _ (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator variable)


**Array vs Slice**  

slice = indexable and have a length. Unlike arrays this length is allowed to change.  

array
```
var x [5]float64
```
slice (not set length)
```
var x []float64
```