**How to run?**

`go run main.go` 

หลักการของคำสั่งนี้คือมีการ compile จากนั้นไปเก็บไว้ที่ temp directory แล้วรันโปรแกรม

package main ~> method main

**Package**

```
package main
```

เรียกว่า package declaration ทุกการเขียนจะต้องมีการประกาศตัวนี้ เป็นการทำให้สามารถ organizing และ reusing code ในภาษา go

รูปแบบโปรแกรมของ Go มี 2 ประเภท
1. executables โปรแกรมที่เราสามารถรันได้จาก command ถ้าใน Windows คือ .exe นั่นเอง
2. libraries เป็นการเขียนโปรแกรมที่เรานำมาทำเป็น package ไว้ เพื่อให้ตอนที่เราเขียนโปรแกรมอื่นๆ มาเรียกใช้งานต่อได้

* White space = Newlines, spaces and tabs (because you can't see them)

**Import**
```
import "fmt"
```

`import` เป็นคำสั่งที่ทำให้เราสามารถนำเอา code จาก package อื่น มาใช้ที่นี้ จากตัวอย่างคือ fmt package (shorthand for format)

" = double quotes (string literal ซึ่งคือ type ของ expression) 

strings = a sequence of characters (letters, numbers, symbols, ...) of a definite length.

* The " character itself is not part of the string

**Comment**

`//` is known as a comment. Comments are ignored by the Go compiler and are there for your own sake (or whoever picks up the source code for your program).

Go supports two diffent styles of comments: `//` comments in which all the text between the `//` and the end of the line is part of the comment and `/* */` comments where everything between the `*`s is part of the comment. (And may include multiple lines)

**Function declaration**
```go
func main() {
	fmt.Println("Hello world :)")
}
```
`func functionName` 

The name `main` is special because it's the function that gets called when you execute the program.

**Body**
```
fmt.Println("Hello world :)")
```

`fmt` package called `Println` (that's the fmt.Println piece, `Println` means Print Line). Then we create a new string that contains Hello World and invoke (also known as call or execute) that function with the string as the first and only argument.

**Document**
```
godoc fmt Println
```
`godoc` แสดงออกมาว่าใน package นั้นมี method อะไรให้ใช้บ้าง

