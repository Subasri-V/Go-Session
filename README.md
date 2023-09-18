// variable 

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


// functions

// 1.  How are functions declared in Go?

// import "fmt"

// func greet(name string) {
//     fmt.Println("Hello,", name)
// }

// func main() {
//     greet("Alice")
// }

// 2. What is a function signature, and what does it include?

// A function signature includes the function's name, parameter types, and return types. It describes the function's interface.

// package main

// import "fmt"

// func add(a, b int) int {
//     return a + b
// }

// func main() {
//     result := add(10, 20)
//     fmt.Println("Result of addition:", result)
// }

// 3. How do you pass arguments to a function in Go?

// package main

// import "fmt"

// func greet(name string) {
//     fmt.Println("Hello,", name)
// }

// func main() {
//     greet("Alice")
// }

// 4.  What is a return statement in a function, and why is it important?

// package main

// import "fmt"

// func add(a, b int) int {
//     return a + b
// }

// func main() {
//     result := add(10, 20)
//     fmt.Println("Result of addition:", result)
// }

// 5. Can a Go function return multiple values? If yes, how is it done?

// package main

// import "fmt"

// func divide(a, b int) (int, int) {
//     quotient := a / b
//     remainder := a % b
//     return quotient, remainder
// }

// func main() {
//     q, r := divide(10, 3)
//     fmt.Println("Quotient:", q)
//     fmt.Println("Remainder:", r)
// }

// 6. Explain the concept of variadic functions in Go.

// Variadic functions in Go allow you to pass a variable number of arguments to a function.
// They are indicated by an ellipsis ... before the parameter type in the function signature.

// package main

// import "fmt"

// func sum(nums ...int) int {
//     total := 0
//     for _, num := range nums {
//         total += num
//     }
//     return total
// }

// func main() {
//     result := sum(1, 2, 3, 4, 5)
//     fmt.Println("Sum:", result)
// }

// 7. What is a function closure in Go, and why is it useful?

// package main

// import "fmt"

// func main() {
//     // Step 1: Define a function that returns a closure.
//     closureFunc := createClosure()

//     // Step 4: Call the closure and store the result.
//     result := closureFunc()

//     // Step 7: Print the result.
//     fmt.Println(result) // Output: 11
// }

// // Step 2: Define a function that returns a closure.
// func createClosure() func() int {
//     // Step 3: Declare a variable within the enclosing function's scope.
//     x := 10

//     // Step 5: Define and return a closure function.
//     // This closure captures 'x' from the surrounding scope.
//     return func() int {
//         // Step 6: Access the captured variable 'x'.
//         return x + 1
//     }
// }

// 8. How do you define and call recursive functions in Go?

// package main

// import "fmt"

// func factorial(n int) int {
//     if n <= 1 {
//         return 1 // Base case
//     }
//     return n * factorial(n-1)
// }

// func main() {
//     result := factorial(5)
//     fmt.Println("Factorial of 5:", result)
// }

// 9. What is the difference between call by value and call by reference in function arguments?

// package main

// import "fmt"

// func modifyValue(val int) {
//     val = 10
// }

// func modifySlice(slice []int) {
//     slice[0] = 10
// }

// func main() {
//     // Call by value for primitive type
//     x := 5
//     modifyValue(x)
//     fmt.Println("Value after modifyValue:", x) // Output: Value after modifyValue: 5

//     // Call by reference (value) for slice
//     numbers := []int{1, 2, 3}
//     modifySlice(numbers)
//     fmt.Println("Slice after modifySlice:", numbers) // Output: Slice after modifySlice: [10 2 3]
// }

// 10. How can you return a function from another function in Go?

// package main

// import "fmt"

// func adder(a int) func(int) int {
//     return func(b int) int {
//         return a + b
//     }
// }

// func main() {
//     addFive := adder(5)
//     result := addFive(10)
//     fmt.Println("Result:", result) // Output: 15
// }
