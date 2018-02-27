**Concurrency**  
more than one task simultaneously is known as concurrency

**Problems**  
1. What's the difference between a method and a function?  
Ref: https://stackoverflow.com/questions/8263546/whats-the-difference-of-functions-and-methods-in-go  

No, that's not correct. There are just a couple of functions from the builtin package which are always available. Everything else needs to be imported.  

The term "method" came up with object-oriented programming. In an OOP language (like C++ for example) you can define a "class" which encapsulates data and functions which belong together. Those functions inside a class are called "methods" and you need an instance of that class to call such a method.  

In Go, the terminology is basically the same, although Go isn't an OOP language in the classical meaning. In Go, a function which takes a receiver is usually called a method (probably just because people are still used to the terminology of OOP).  

So, for example:  
```go
func MyFunction(a, b int) int {
  return a + b
}
// Usage:
// MyFunction(1, 2)
```
but
```go
type MyInteger int
func (a MyInteger) MyMethod(b int) int {
  return a + b
}
// Usage:
// var x MyInteger = 1
// x.MyMethod(2)
```


2. Why would you use an embedded anonymous field instead of a normal named field?
would like to inherit from other struct   


