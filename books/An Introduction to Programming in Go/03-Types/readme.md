**Type VS Token**

*Example*  
Max = name of Dog  
Max = Token (a particular instance or member)  
Dog = Type (the general concept)   
All dogs have 4 legs, Max is a dog, therefore Max has 4 legs.  
~> All strings have a length, x is a string, therefore x has a length.  

mathematics -> sets  
R (the set of all real numbers) or N (the set of all natural numbers)  
Each member of these sets shares properties with all the other members of the set.  
*“for all natural numbers a, b, and c, a + (b + c) = (a + b) + c and a × (b × c) = (a × b) × c.”*   

**Types**

1. Numbers  
1.1. Integer  
	Integers – like their mathematical counterpart – are numbers without a decimal component. (..., -3, -2, -1, 0, 1, ...) Unlike the base-10 decimal system we use to represent numbers, computers use a base-2 binary system uint8, uint16, uint32, uint64, int8, int16, int32 and int64. 8, 16, 32 and 64 tell us how many bits each of the types use. uint means “unsigned integer” while int means “signed integer”.  
	byte = unit8   
	rune = int32  
	There are also 3 machine de- pendent integer types: uint, int and uintptr.  
	Go's byte data type is often used in the definition of other types. There are also 3 machine de- pendent integer types: uint, int and uintptr.  
1.2. Float  
	computing 1.01 - 0.99 results in 0.020000000000000018 – A number extremely close to what we would expect, but not exactly the same.  
	Like integers floating point numbers have a cer- tain size (32 bit or 64 bit). Using a larger sized floating point number increases it's precision. (how many digits it can represent)  
	In addition to numbers there are several other values which can be represented: “not a number” (NaN, for things like 0/0) and positive and negative infinity. (+∞ and −∞) 
	float32 / float64 -> complex numbers (numbers with imaginary parts): complex64 and complex128  

2. String  
"Hello world"  
\`Hello World  
new line  
Hello World\` 

3. Boolean  
named after [George Boole](https://en.wikipedia.org/wiki/George_Boole)  

**Problems**
1. How are integers stored on acomputer?  
	binary digit  
2. what's the largest 8 digit number?  
	2^8 - 1 = 255  
3. 3-problem.go
4. What is a string? How do you find its length?  
	array of a bytes, len
