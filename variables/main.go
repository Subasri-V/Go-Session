// 1.What is a variable in Go, and how is it declared?

package main

// import "fmt"

// func main() {
// 	var age int // Declaring an integer variable named "age"
// 	age = 30
// 	fmt.Println("Age:", age)
// }

// 2.Explain the difference between a short variable declaration and a regular variable declaration.

// A short variable declaration (:=) is used for declaring and initializing variables within a function or block scope,
// while a regular variable declaration (var) is used for package-level or global variables.

// package main

// import "fmt"

// func main() {
// 	name := "John"             // Short variable declaration
// 	var surname string = "Doe" // Regular variable declaration
// 	fmt.Println("Name:", name)
// 	fmt.Println("Surname:", surname)
// }

// 3.What is the zero value of a variable in Go?

// package main

// import "fmt"

// func main() {
//     var count int     // Zero value of an integer variable is 0
//     var price float64 // Zero value of a float64 variable is 0.0
//     var name string   // Zero value of a string variable is an empty string
//     fmt.Println("Count:", count)
//     fmt.Println("Price:", price)
//     fmt.Println("Name:", name)
// }

// 4. How do you initialize a variable during declaration?

// package main

// import "fmt"

// func main() {
//     var temperature float64 = 25.5 // Initializing a float64 variable
//     fmt.Println("Temperature:", temperature)
// }

// 5.What is the scope of a variable in Go?

// Variables declared within a function have local scope, while those declared at the package level have a wider scope.

// package main

// import "fmt"

// var globalVar = 42 // This variable has package-level scope

// func main() {
//     var localVar = 10 // This variable has local scope within the main function
//     fmt.Println("Local variable:", localVar)
//     fmt.Println("Global variable:", globalVar)
// }

// 6. What is shadowing in Go, and why should it be avoided?

// Shadowing occurs when a variable declared in an inner scope has the same name as a variable in an outer scope, effectively hiding the outer variable within that inner scope.
// Shadowing should be avoided because it can lead to confusion and unexpected behavior.

// package main

// import "fmt"

// func main() {
//     age := 30 // Outer age variable
//     if true {
//         age := 40 // Inner age variable, shadowing the outer one
//         fmt.Println("Inner age:", age)
//     }
//     fmt.Println("Outer age:", age)
// }

// 7. How can you declare and assign multiple variables in a single line?

// package main

// import "fmt"

// func main() {
//     a, b := 10, 20 // Declaring and initializing multiple variables in one line
//     fmt.Println("a:", a)
//     fmt.Println("b:", b)
// }

// 8. What is the difference between var, const, and := when declaring variables?

// package main

// import "fmt"

// func main() {
//     var x int // Variable declaration with 'var'
//     const pi = 3.14 // Constant declaration with 'const'
//     name := "Alice" // Short variable declaration with ':='
//     fmt.Println("x:", x)
//     fmt.Println("pi:", pi)
//     fmt.Println("name:", name)
// }

// 9. Can you change the data type of a variable once it's declared?

// yes or no

// No, you cannot change the data type of a variable once it's declared. Go is a statically typed language, and variable data types are determined at compile time.

// 10 . Explain the concept of naming conventions for variables in Go.

// package main

// import "fmt"

// func main() {
//     myVariableName := "example" // Good: myVariableName
//     MyVariable := "example"     // Bad: MyVariable (starts with an uppercase letter)
//     my_variable_name := "example" // Bad: my_variable_name (uses underscores)
//     fmt.Println("Variable name:", myVariableName)
// }
