package main

// Question 1: What is an anonymous function in Go, and how is it declared?

// import "fmt"

// func main() {
// 	fmt.Println("Anonymous Functions in Go:")
// 	add := func(a, b int) int {
// 		return a + b
// 	}
// 	result := add(3, 5)
// 	fmt.Println("Result:", result)
// }
// Options:

// a. Anonymous functions in Go are functions with no body.
// b. Anonymous functions in Go are functions with no parameters.
// c. Anonymous functions in Go are functions without a name and can be declared inline.
// d. Anonymous functions in Go are functions that can only be used once.
// Answer: c. Anonymous functions in Go are functions without a name and can be declared inline.

// Explanation: In Go, anonymous functions are functions without a name that can be declared and used inline within the code. They are often used for short, specific tasks and can capture variables from their surrounding lexical scope.

// Question 2: Explain the concept of closures in relation to anonymous functions.

// import "fmt"

// func main() {
// 	fmt.Println("Closures in Go:")

// 	x := 10
// 	closure := func() {
// 		fmt.Println("x from closure:", x)
// 	}
// 	closure()

// 	x = 20

// 	closure()
// }
// Options:

// a. Closures in Go refer to the automatic closing of anonymous functions after execution.
// b. Closures are a way to define multiple anonymous functions within a single function.
// c. Closures allow anonymous functions to capture and remember the variables from their surrounding lexical scope.
// d. Closures are a feature specific to named functions and cannot be used with anonymous functions.
// Answer: c. Closures allow anonymous functions to capture and remember the variables from their surrounding lexical scope.

// Explanation: Closures in Go allow anonymous functions to capture and access variables from their surrounding lexical scope, even after the surrounding function has finished execution.

// Question 3: How can you pass anonymous functions as arguments to other functions?

// import "fmt"

// func main() {
// 	fmt.Println("Passing Anonymous Functions in Go:")

// 	greeter := func(name string) {
// 		fmt.Println("Hello,", name)
// 	}

// 	greetSomeone(greeter)
// }

// func greetSomeone(greeter func(string)) {
// 	greeter("Alice")
// }
// Options:

// a. Anonymous functions cannot be passed as arguments to other functions in Go.
// b. You can pass an anonymous function as an argument by directly invoking it within the function call.
// c. You can pass an anonymous function as an argument by defining it as a parameter with a function type.
// d. You can pass an anonymous function as an argument by using the return keyword.
// Answer: c. You can pass an anonymous function as an argument by defining it as a parameter with a function type.

// Explanation: In Go, you can pass anonymous functions as arguments to other functions by defining the parameter with a function type that matches the signature of the anonymous function.

// Question 4: Provide an example of using an anonymous function as a callback.

// import "fmt"

// func main() {
// 	fmt.Println("Using Anonymous Functions as Callbacks in Go:")
// 	numbers := []int{1, 2, 3, 4, 5}
// 	processNumbers(numbers, func(num int) {
// 		fmt.Println("Processed number:", num)
// 	})
// }

// func processNumbers(numbers []int, callback func(int)) {
// 	for _, num := range numbers {
// 		callback(num)
// 	}
// }
// Options:

// a. Anonymous functions cannot be used as callbacks in Go.
// b. Here is an example of using an anonymous function as a callback in Go.
// c. Callbacks can only be defined as named functions in Go.
// d. Callbacks are not a concept in Go programming.
// Answer: b. Here is an example of using an anonymous function as a callback in Go.

// Explanation: Anonymous functions can be used as callbacks in Go to perform specific actions for each element in a collection, such as processing numbers in this example.

// Question 5: What is the advantage of using anonymous functions over named functions in certain situations?

// import "fmt"

// func main() {
// 	fmt.Println("Advantages of Anonymous Functions in Go:")

// 	double := func(x int) int {
// 		return x * 2
// 	}

// 	result := double(5)
// 	fmt.Println("Result:", result)
// }
// Options:

// a. Anonymous functions are always less efficient than named functions in Go.
// b. Anonymous functions cannot access variables from their surrounding scope, making them less flexible.
// c. Anonymous functions allow you to define and use functions without cluttering the global scope and can capture variables from their enclosing scope.
// d. Anonymous functions are only used for debugging purposes and should be avoided in production code.
// Answer: c. Anonymous functions allow you to define and use functions without cluttering the global scope and can capture variables from their enclosing scope.

// Explanation: Anonymous functions are advantageous in scenarios where you want to define and use functions without polluting the global scope, and they can capture variables from their surrounding lexical scope, making them versatile for various use cases.

// Question 6:Predict the output

// package main
// import "fmt"
// func delayedGreeting(name string) func() {
//   return func() {
//     fmt.Println("Hello, " + name + "!")
//   }
// }
// func main() {
//   greet := delayedGreeting("Alice")
//   defer greet()
//   fmt.Println("Welcome!")
// }

// a. Hello Alice welcome!
// b. Hello Alice
//    Welcome!
// c. Wecome!
//    Hello Alice !
// d. Welcome! Hello Alice !
// Answer : c. Wecome!
//             Hello Alice !


// Question 7:Predict the output
// package main
// import "fmt"
// func adder() func(int) int {
//   sum := 0
//   return func(x int) int {
//     sum += x
//     return sum
//   }
// }
// func main() {
//   add := adder()
//   fmt.Println(add(1))
//   fmt.Println(add(2)) 
//   fmt.Println(add(3)) 
// }

// a. 1 3 6
// b. 1
//    3  
//    6
// c. 1
//    6
//    3
// d.1 6 3

// Answer : b. 1
//             3  
//             6


