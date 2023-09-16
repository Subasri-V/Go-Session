// 1.  How are functions declared in Go?

package main

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
