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

   ## Pointers 

   Pointers in Go are easy and fun to learn. Some Go programming tasks are performed more easily with pointers, and other tasks, such as call by reference, cannot be performed without using pointers. So it becomes necessary to learn pointers to become a perfect Go programmer.

   As you know, every variable is a memory location and every memory location has its address defined which can be accessed using ampersand (&) operator, which denotes an address in memory. Consider the following example, which will print the address of the variables defined −

   ## What Are Pointers?
   A pointer is a variable whose value is the address of another variable, i.e., direct address of the memory location. Like any variable or constant, you must declare a pointer before you can use it to store any variable address. The general form of a pointer variable declaration is −

   `var var_name *var-type`
   `var ip *int /* pointer to an integer */`
   `var fp *float32 /* pointer to a float */`

 ## Structures
   Structure is another user-defined data type available in Go programming, which allows you to combine data items of different kinds.
   
   
   `type struct_variable_type struct {`
      `member definition;`
      `member definition;`
      `...`
      `member definition;`
   `}`
   
   Once a structure type is defined, it can be used to declare variables of that type using the following syntax.

   `variable_name := structure_variable_type {value1, value2...valuen}`
