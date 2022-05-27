## Go Language Learning Resource 
<p> In This Repository : I Update Topic That I Will learn </p>

## Hello World  
 `go run file name`  : Run A Specific File
 `go run .`

 ## Hello World Example
A Go program basically consists of the following parts −

1) Package Declaration
2) Import Packages
3) Functions
4) Variables
5) Statements and Expressions
6) Comments

## Variable Declaration In Short and Declaractive Way

Short `name := "Bangladesh"`

Declaractive `var name string;`
`name = "Bangladesh"`

## Function Arguments
If a function is to use arguments, it must declare variables that accept the values of the arguments. These variables are called the formal parameters of the function.

The formal parameters behave like other local variables inside the function and are created upon entry into the function and destroyed upon exit.

While calling a function, there are two ways that arguments can be passed to a function −

Sr.No	Call Type & Description
1	Call by value
This method copies the actual value of an argument into the formal parameter of the function. In this case, changes made to the parameter inside the function have no effect on the argument.

2	Call by reference
This method copies the address of an argument into the formal parameter. Inside the function, the address is used to access the actual argument used in the call. This means that changes made to the parameter affect the argument.