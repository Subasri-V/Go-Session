package main

// Functions:

// Question 1: How are functions declared in Go?

// import "fmt"

// func greet(name string) string {
//     return "Hello, " + name + "!"
// }

// func main() {
//     fmt.Println("Function Declaration in Go:")
//     message := greet("Alice")
//     fmt.Println(message)
// }
// Options:

// a. Functions are declared using the func keyword followed by the function name and parameters.
// b. Functions are declared using the function keyword in Go.
// c. Functions are declared using the def keyword followed by the function name and parameters.
// d. Functions are declared using the fn keyword in Go.
// Answer: a. Functions are declared using the func keyword followed by the function name and parameters.

// Question 2: What is a function signature, and what does it include?

// import "fmt"

// func calculateSum(a, b int) int {
//     return a + b
// }

// func main() {
//     fmt.Println("Function Signature in Go:")
//     result := calculateSum(5, 3)
//     fmt.Println("Sum:", result)
// }
// Options:

// a. A function signature includes only the function name.
// b. A function signature includes the function name, return type, and parameter list.
// c. A function signature includes comments and documentation.
// d. A function signature includes the function name and the package name.
// Answer: b. A function signature includes the function name, return type, and parameter list.

// Question 3: How do you pass arguments to a function in Go?

// import "fmt"

// func greet(name string) string {
//     return "Hello, " + name + "!"
// }

// func main() {
//     fmt.Println("Passing Arguments to a Function in Go:")
//     message := greet("Bob")
//     fmt.Println(message)
// }
// Options:

// a. Arguments are passed using parentheses without specifying parameter names.
// b. Arguments are passed using curly braces.
// c. Arguments are passed using square brackets.
// d. Arguments are not allowed in Go functions.
// Answer: a. Arguments are passed using parentheses without specifying parameter names.

// Question 4: What is a return statement in a function, and why is it important?

// import "fmt"

// func add(a, b int) int {
//     result := a + b
//     return result // Return statement
// }

// func main() {
//     fmt.Println("Return Statement in a Function:")
//     sum := add(7, 3)
//     fmt.Println("Sum:", sum)
// }
// Options:

// a. A return statement is used to exit the function, and it's not necessary in Go.
// b. A return statement is used to specify the function's name.
// c. A return statement is used to send a value back to the caller, and it's important for providing results.
// d. A return statement is used to print output to the console.
// Answer: c. A return statement is used to send a value back to the caller, and it's important for providing results.

// Question 5: Can a Go function return multiple values? If yes, how is it done?

// import "fmt"

// func divide(a, b int) (int, error) {
//     if b == 0 {
//         return 0, fmt.Errorf("division by zero")
//     }
//     return a / b, nil
// }

// func main() {
//     fmt.Println("Returning Multiple Values from a Function:")
//     result, err := divide(10, 2)
//     if err != nil {
//         fmt.Println("Error:", err)
//     } else {
//         fmt.Println("Result:", result)
//     }
// }
// Options:

// a. Go functions cannot return multiple values.
// b. Yes, Go functions can return multiple values by specifying them within parentheses.
// c. Go functions return multiple values using square brackets.
// d. Multiple return values require the use of a special multireturn keyword.
// Answer: b. Yes, Go functions can return multiple values by specifying them within parentheses.

// Question 6: Explain the concept of variadic functions in Go.

// import "fmt"

// func average(numbers ...float64) float64 {
//     total := 0.0
//     count := 0
//     for _, num := range numbers {
//         total += num
//         count++
//     }
//     if count == 0 {
//         return 0.0
//     }
//     return total / float64(count)
// }

// func main() {
//     fmt.Println("Variadic Functions in Go:")
//     avg := average(2.0, 4.0, 6.0, 8.0)
//     fmt.Printf("Average: %.2f\n", avg)
// }
// Options:

// a. Variadic functions are functions that have a fixed number of parameters.
// b. Variadic functions are functions that accept a variable number of arguments of the same type.
// c. Variadic functions are functions that can return multiple values.
// d. Variadic functions are functions that cannot have any parameters.
// Answer: b. Variadic functions are functions that accept a variable number of arguments of the same type.

// Question 7: What is a function closure in Go, and why is it useful?

// import "fmt"

// func increment() func() int {
//     count := 0
//     return func() int {
//         count++
//         return count
//     }
// }

// func main() {
//     fmt.Println("Function Closure in Go:")
//     incrementFn := increment()

//     fmt.Println("Increment 1:", incrementFn())
//     fmt.Println("Increment 2:", incrementFn())
//     fmt.Println("Increment 3:", incrementFn())
// }
// Options:

// a. Function closure is a way to declare functions within other functions, and it's rarely used in Go.
// b. Function closure is a technique to hide the implementation details of a function.
// c. Function closure is a way to encapsulate a function and its local variables, allowing them to persist across multiple calls.
// d. Function closure is a way to create anonymous functions in Go.
// Answer: c. Function closure is a way to encapsulate a function and its local variables, allowing them to persist across multiple calls.

// Question 8: How do you define and call recursive functions in Go?

// import "fmt"

// func factorial(n int) int {
//     if n <= 1 {
//         return 1
//     }
//     return n * factorial(n-1) // Recursive call
// }

// func main() {
//     fmt.Println("Recursive Functions in Go:")
//     result := factorial(5)
//     fmt.Println("Factorial of 5:", result)
// }
// Options:

// a. Recursive functions are not allowed in Go due to potential stack overflow issues.
// b. Recursive functions are defined using the rec keyword, and they can only call themselves once.
// c. Recursive functions are defined by specifying the recursive modifier, and they can call other functions recursively.
// d. Recursive functions are defined by calling themselves within their own definition.
// Answer: d. Recursive functions are defined by calling themselves within their own definition.

// Question 9: What is the difference between call by value and call by reference in function arguments?

// import "fmt"

// func modifyValue(val int) {
//     val = 10
// }

// func modifySlice(slice []int) {
//     slice[0] = 10
// }

// func main() {
//     fmt.Println("Call by Value vs. Call by Reference:")
//     x := 5
//     modifyValue(x)
//     fmt.Println("Value after modifyValue:", x)

//     nums := []int{1, 2, 3}
//     modifySlice(nums)
//     fmt.Println("Slice after modifySlice:", nums)
// }
// Options:

// a. Call by value means passing a copy of the original value, and call by reference means passing a reference to the original value.
// b. Call by value means passing a reference to the original value, and call by reference means passing a copy of the original value.
// c. Call by value is a Go-specific feature, and call by reference is a general programming concept.
// d. Call by value and call by reference are the same in Go.
// Answer: a. Call by value means passing a copy of the original value, and call by reference means passing a reference to the original value.

// Question 10: How can you return a function from another function in Go?

// import "fmt"

// func getMultiplier(factor int) func(int) int {
//     return func(x int) int {
//         return x * factor
//     }
// }

// func main() {
//     fmt.Println("Returning a Function from Another Function:")
//     timesTwo := getMultiplier(2)
//     timesThree := getMultiplier(3)

//     fmt.Println("2 times 5:", timesTwo(5))
//     fmt.Println("3 times 5:", timesThree(5))
// }
// Options:

// a. Returning a function from another function is not possible in Go.
// b. You can only return predefined functions, not user-defined functions.
// c. You can return a function by defining a new type for it.
// d. You can return a function as a value, allowing you to call it later.
// Answer: d. You can return a function as a value, allowing you to call it later.
