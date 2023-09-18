package main

// Variables:

// Question 1: What is a variable in Go, and how is it declared?

// import "fmt"

// func main() {
//     var name string = "John"
//     fmt.Println("A variable in Go:")
//     fmt.Println("Name:", name)
// }
// Options:

// a. A variable is a fixed value in Go, and it is declared using the val keyword.
// b. A variable is a named storage location for data in Go, and it is declared using the var keyword.
// c. A variable is a constant value in Go, and it is declared using the const keyword.
// d. A variable is a special data type in Go that doesn't require declaration.
// Answer: b. A variable is a named storage location for data in Go, and it is declared using the var keyword.

// Question 2: Explain the difference between a short variable declaration and a regular variable declaration.

// import "fmt"

// func main() {
//     name := "Alice" // Short variable declaration
//     var age int     // Regular variable declaration

//     fmt.Println("Short vs. Regular Variable Declaration:")
//     fmt.Println("Name:", name)
//     fmt.Println("Age:", age)
// }
// Options:

// a. There is no difference between short and regular variable declarations; they are interchangeable.
// b. Short variable declarations use the var keyword, while regular variable declarations use :=.
// c. Short variable declarations are for local variables, while regular variable declarations are for global variables.
// d. Short variable declarations automatically initialize variables, while regular declarations do not.
// Answer: d. Short variable declarations automatically initialize variables, while regular declarations do not.

// Question 3: What is the zero value of a variable in Go?

// import "fmt"

// func main() {
//     var count int
//     var price float64
//     var name string
//     var active bool

//     fmt.Println("Zero Value of Variables in Go:")
//     fmt.Println("Count:", count)
//     fmt.Println("Price:", price)
//     fmt.Println("Name:", name)
//     fmt.Println("Active:", active)
// }
// Options:

// a. The zero value of a variable is the initial value that you assign to it.
// b. The zero value of a variable is the value 0 for numeric types, an empty string for strings, and false for booleans.
// c. The zero value of a variable is a special value that represents null.
// d. The zero value of a variable is an error condition in Go.
// Answer: b. The zero value of a variable is the value 0 for numeric types, an empty string for strings, and false for booleans.

// Question 4: How do you initialize a variable during declaration?

// import "fmt"

// func main() {
//     var age int = 30 // Variable initialization
//     fmt.Println("Initializing a Variable:")
//     fmt.Println("Age:", age)
// }
// Options:

// a. You can only initialize variables with the := operator.
// b. Variable initialization is not allowed in Go.
// c. You can initialize a variable using the = operator during declaration.
// d. Variable initialization can only be done in a separate assignment statement.
// Answer: c. You can initialize a variable using the = operator during declaration.

// Question 5: What is the scope of a variable in Go?

// import "fmt"

// var globalVariable = "I'm a global variable" // Global scope

// func main() {
//     var localVariable = "I'm a local variable" // Local scope
//     fmt.Println("Variable Scopes in Go:")
//     fmt.Println(globalVariable)
//     fmt.Println(localVariable)
// }

// func anotherFunction() {
//     // Cannot access localVariable here
//     fmt.Println("Accessing Global Variable from Another Function:")
//     fmt.Println(globalVariable)
// }
// Options:

// a. Variables in Go have only global scope.
// b. Variables in Go have only local scope.
// c. Variables in Go have both global and local scope.
// d. Variable scope depends on the data type.
// Answer: c. Variables in Go have both global and local scope.

// Question 6: What is shadowing in Go, and why should it be avoided?

// import "fmt"

// var count int = 10 // Global variable

// func main() {
//     var count int = 5 // Local variable shadowing global variable
//     fmt.Println("Shadowing in Go:")
//     fmt.Println("Local Count:", count)
// }

// func anotherFunction() {
//     fmt.Println("Global Count:", count) // Accessing the global variable
// }
// Options:

// a. Shadowing is a feature that allows you to define the same variable multiple times for clarity.
// b. Shadowing is a technique to hide global variables and prioritize local variables.
// c. Shadowing is a mistake where a local variable with the same name as a global variable hides the global variable within a specific scope.
// d. Shadowing is a recommended practice for improving code readability.
// Answer: c. Shadowing is a mistake where a local variable with the same name as a global variable hides the global variable within a specific scope.

// Question 7: How can you declare and assign multiple variables in a single line?

// import "fmt"

// func main() {
//     var a, b, c int // Declaration
//     a, b, c = 1, 2, 3 // Assignment

//     fmt.Println("Declaring and Assigning Multiple Variables:")
//     fmt.Println("a:", a)
//     fmt.Println("b:", b)
//     fmt.Println("c:", c)
// }
// Options:

// a. You cannot declare and assign multiple variables in a single line in Go.
// b. Use the var keyword followed by a list of variables and assign values using =.
// c. Use the := operator to declare and assign multiple variables simultaneously.
// d. You can only assign values to variables one at a time.
// Answer: b. Use the var keyword followed by a list of variables and assign values using =.

// Question 8: What is the difference between var, const, and := when declaring variables?

// import "fmt"

// func main() {
//     var a int = 10 // Using var
//     const b = 20   // Using const
//     c := 30         // Using :=

//     fmt.Println("Variable Declaration in Go:")
//     fmt.Println("a:", a)
//     fmt.Println("b:", b)
//     fmt.Println("c:", c)
// }
// Options:

// a. var is used for declaring variables with initial values, const is used for constants, and := is used for short variable declarations.
// b. var is used for constants, const is used for declaring variables with initial values, and := is used for long variable declarations.
// c. var is used for constants, const is used for short variable declarations, and := is used for declaring variables with initial values.
// d. var and const are the same, and := is used for declaring variables without initial values.
// Answer: a. var is used for declaring variables with initial values, const is used for constants, and := is used for short variable declarations.

// Question 9: Can you change the data type of a variable once it's declared?

// import "fmt"

// func main() {
//     var age int = 30
//     age = 35 // Changing the value (not the data type)

//     fmt.Println("Changing the Data Type of a Variable:")
//     fmt.Println("Age:", age)
// }
// Options:

// a. Yes, you can change the data type of a variable at any time.
// b. No, once a variable is declared with a specific data type, you cannot change it.
// c. You can change the data type by using type casting.
// d. Changing the data type requires redeclaration of the variable.
// Answer: b. No, once a variable is declared with a specific data type, you cannot change it.

// Question 10: Explain the concept of naming conventions for variables in Go.

// import "fmt"

// func main() {
//     camelCaseVariable := "camelCase"
//     snake_case_variable := "snake_case"
//     PascalCaseVariable := "PascalCase"

//     fmt.Println("Naming Conventions for Variables in Go:")
//     fmt.Println("Camel Case:", camelCaseVariable)
//     fmt.Println("Snake Case:", snake_case_variable)
//     fmt.Println("Pascal Case:", PascalCaseVariable)
// }
// Options:

// a. Go has no specific naming conventions for variables; you can use any style you prefer.
// b. Go uses camel case for variables, snake case for constants, and Pascal case for types.
// c. Go uses snake case for variables, camel case for constants, and Pascal case for types.
// d. Go requires all variables to be named in uppercase letters.
// Answer: b. Go uses camel case for variables, snake case for constants, and Pascal case for types.
