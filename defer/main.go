package main

// Defer:

// Question 1: What is the purpose of the defer statement in Go?

// import "fmt"

// func main() {
// 	fmt.Println("Defer Statement in Go:")
// 	defer fmt.Println("Deferred statement executed")
// 	fmt.Println("Regular statement")
// }
// Options:

// a. Defer is used to skip the execution of a statement.
// b. Defer is used to execute a statement immediately.
// c. Defer is used to delay the execution of a statement until the surrounding function returns.
// d. Defer is used to exit the program.
// Answer: c. Defer is used to delay the execution of a statement until the surrounding function returns.

// Question 2: How does the defer statement work in terms of execution order?

// import "fmt"

// func main() {
// 	fmt.Println("Execution Order of Defer Statements:")
// 	defer fmt.Println("Deferred statement 1")
// 	fmt.Println("Regular statement 1")
// 	defer fmt.Println("Deferred statement 2")
// 	fmt.Println("Regular statement 2")
// }
// Options:

// a. Deferred statements are executed before regular statements.
// b. Deferred statements are executed after regular statements.
// c. Deferred statements are executed in the reverse order of their appearance.
// d. The execution order of deferred statements is unpredictable.
// Answer: c. Deferred statements are executed in the reverse order of their appearance.

// Question 3: Give an example of a practical use case for the defer statement.

// import "fmt"

// func main() {
// 	fmt.Println("Practical Use of Defer in Go:")
// 	file := openFile("example.txt")
// 	defer closeFile(file)
// 	readFile(file)
// }

// func openFile(filename string) int {
// 	fmt.Printf("Opening file: %s\n", filename)
// 	// Simulate opening the file and returning a file descriptor
// 	return 123
// }

// func closeFile(fd int) {
// 	fmt.Printf("Closing file with fd: %d\n", fd)
// 	// Simulate closing the file
// }

// func readFile(fd int) {
// 	fmt.Printf("Reading from file with fd: %d\n", fd)
// 	// Simulate reading from the file
// }
// Options:

// a. Defer is used to open files in Go programs.
// b. Defer is used to prevent the execution of functions.
// c. Defer is used to ensure that resources are released after they are no longer needed, such as closing files or network connections.
// d. Defer is used to execute functions before their regular order.
// Answer: c. Defer is used to ensure that resources are released after they are no longer needed, such as closing files or network connections.

// Question 4: What happens when multiple defer statements are used in a function?

// import "fmt"

// func main() {
// 	fmt.Println("Multiple Defer Statements in a Function:")
// 	defer fmt.Println("Deferred statement 1")
// 	defer fmt.Println("Deferred statement 2")
// 	defer fmt.Println("Deferred statement 3")
// 	fmt.Println("Regular statement")
// }
// Options:

// a. Only the first defer statement is executed, and the rest are ignored.
// b. All defer statements are executed in the order they are encountered.
// c. Only the last defer statement is executed, and the others are overridden.
// d. The order of execution of defer statements is random.
// Answer: c. Only the last defer statement is executed, and the others are overridden.

// Question 5: Can you pass arguments to a deferred function call?

// import "fmt"

// func main() {
// 	fmt.Println("Passing Arguments to a Deferred Function:")
// 	value := 42
// 	defer fmt.Println("Value inside deferred function:", value)
// 	value = 0
// }
// Options:

// a. Yes, you can pass arguments to a deferred function call, and the arguments retain their values when the deferred function is executed.
// b. No, you cannot pass arguments to a deferred function call.
// c. Passing arguments to a deferred function call is optional, and you can choose whether or not to pass them.
// d. Passing arguments to a deferred function call is only allowed for built-in functions.
// Answer: a. Yes, you can pass arguments to a deferred function call, and the arguments retain their values when the deferred function is executed.

// Question 6: Is it possible to change the return value of a function with defer?

// import "fmt"

// func main() {
// 	fmt.Println("Changing Return Value with Defer:")
// 	result := calculate()
// 	fmt.Println("Result:", result)
// }

// func calculate() int {
// 	result := 10
// 	defer func() {
// 		result = 20
// 	}()
// 	return result
// }
// Options:

// a. Yes, you can change the return value of a function using defer.
// b. No, defer cannot modify the return value of a function.
// c. Changing the return value with defer requires the use of special keywords.
// d. Defer can only modify the return value of functions with a specific declaration.
// Answer: a. Yes, you can change the return value of a function using defer.

// Question 7: Explain the key difference between defer and panic/recover in Go.

// import "fmt"

// func main() {
// 	fmt.Println("Defer vs. Panic/Recover in Go:")
// 	defer fmt.Println("Deferred statement")
// 	panic("Something went wrong")
// 	fmt.Println("After panic")
// }

// Options:

// a. Both defer and panic/recover are used for error handling, and they are interchangeable.
// b. Defer is used to handle errors, while panic/recover is used for cleanup operations.
// c. Defer is used for cleanup operations, while panic/recover is used for handling and propagating errors.
// d. Defer and panic/recover have no differences in their usage.
// Answer: c. Defer is used for cleanup operations, while panic/recover is used for handling and propagating errors.

// Question 8: How can you ensure that a defer statement executes even if a panic occurs?

// import "fmt"

// func main() {
// 	fmt.Println("Ensuring Defer Execution During Panic:")
// 	defer fmt.Println("Deferred statement")
// 	panic("Something went wrong")
// 	fmt.Println("After panic")
// }

// Options:

// a. Defer statements are automatically executed during a panic, so no special handling is needed.
// b. You cannot ensure that a defer statement executes during a panic; it depends on the specific circumstances.
// c. You can use the recover function to capture and suppress a panic, allowing the deferred statements to execute.
// d. Defer statements are never executed during a panic.
// Answer: c. You can use the recover function to capture and suppress a panic, allowing the deferred statements to execute.

// Question 9: Can you use defer with functions that have return values?

// import "fmt"

// func main() {
// 	fmt.Println("Using Defer with Functions Returning Values:")
// 	result := calculate()
// 	fmt.Println("Result:", result)
// }

// func calculate() int {
// 	result := 10
// 	defer func() {
// 		result = 20
// 	}()
// 	return result
// }
// Options:

// a. Defer can only be used with functions that have no return values.
// b. Defer can be used with functions that have return values, but it does not affect the return value.
// c. Defer can be used with functions that have return values, and it can change the return value.
// d. Defer cannot be used with functions that have return values.
// Answer: b. Defer can be used with functions that have return values, but it does not affect the return value.

// Question 10: When should you use defer, and when is it better to avoid it?

// import "fmt"

// func main() {
// 	fmt.Println("When to Use Defer in Go:")
// 	defer fmt.Println("Deferred statement 1")
// 	fmt.Println("Regular statement 1")
// 	defer fmt.Println("Deferred statement 2")
// 	fmt.Println("Regular statement 2")
// 	defer fmt.Println("Deferred statement 3")
// 	fmt.Println("Regular statement 3")
// }
// Options:

// a. Defer should be used for all statements to ensure proper execution.
// b. Defer should be avoided as it leads to unpredictable behavior.
// c. Defer should be used for resource cleanup and actions that must occur before a function exits.
// d. Defer should be used only for debugging purposes.
// Answer: c. Defer should be used for resource cleanup and actions that must occur before a function exits.
