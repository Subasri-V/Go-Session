package main

// Type Conversion:

// Question 1: What is type conversion in Go, and why is it necessary?

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	numStr := "42"
// 	num, err := strconv.Atoi(numStr)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Type Conversion in Go:")
// 	fmt.Printf("Converted value: %d\n", num)
// }
// Options:

// a. Type conversion is a way to create new data types in Go.
// b. Type conversion is a technique to automatically convert between different types in Go.
// c. Type conversion is necessary to ensure type safety when working with different data types.
// d. Type conversion is not supported in Go.
// Answer: c. Type conversion is necessary to ensure type safety when working with different data types.

// Question 2: How do you convert between different numeric types in Go?

// import (
// 	"fmt"
// )

// func main() {
// 	var intNum int32 = 42
// 	floatNum := float64(intNum)

// 	fmt.Println("Converting Between Numeric Types in Go:")
// 	fmt.Printf("Converted float: %.2f\n", floatNum)
// }
// Options:

// a. Numeric type conversion is not allowed in Go.
// b. You can use the convert keyword to change numeric types.
// c. You can convert between numeric types by assigning the value to a variable of the desired type.
// d. Numeric type conversion is automatic in Go, and you don't need to do anything.
// Answer: c. You can convert between numeric types by assigning the value to a variable of the desired type.

// Question 3: Explain type conversion between strings and other data types.

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	num := 42
// 	numStr := strconv.Itoa(num)

// 	fmt.Println("Type Conversion Between Strings and Other Types:")
// 	fmt.Printf("Converted string: %s\n", numStr)
// }
// Options:

// a. Type conversion between strings and other types is not supported in Go.
// b. You can use the strconvert package to convert between strings and other types.
// c. Type conversion between strings and other types is done using the ToString function.
// d. You can use functions like strconv.Itoa and strconv.Atoi to convert between strings and other types.
// Answer: d. You can use functions like strconv.Itoa and strconv.Atoi to convert between strings and other types.

// Question 4: What is type assertion, and when is it used?

// import "fmt"

// func printLength(data interface{}) {
// 	length, ok := data.(int)
// 	if ok {
// 		fmt.Printf("Length: %d\n", length)
// 	} else {
// 		fmt.Println("Invalid data type")
// 	}
// }

// func main() {
// 	value := 42
// 	printLength(value)

// 	str := "Hello"
// 	printLength(str)
// }
// Options:

// a. Type assertion is a way to change the data type of a variable.
// b. Type assertion is used to cast any value to the interface{} type.
// c. Type assertion is a technique to check and extract the underlying value and data type of an interface variable.
// d. Type assertion is not used in Go because Go has a static type system.
// Answer: c. Type assertion is a technique to check and extract the underlying value and data type of an interface variable.

// Question 5: How can you check if a type conversion is valid in Go?

// import (
// 	"fmt"
// )

// func main() {
// 	var num interface{} = 42
// 	str, ok := num.(string)

// 	fmt.Println("Checking Valid Type Conversion in Go:")
// 	if ok {
// 		fmt.Printf("Converted string: %s\n", str)
// 	} else {
// 		fmt.Println("Invalid type conversion")
// 	}
// }
// Options:

// a. You can check the validity of type conversion using the valid keyword in Go.
// b. Type conversion is always valid in Go, so there's no need to check.
// c. You can use a type assertion along with a boolean flag to check the validity of type conversion.
// d. Go does not provide a way to check the validity of type conversion.
// Answer: c. You can use a type assertion along with a boolean flag to check the validity of type conversion.

// Question 6: What are the potential risks of type conversion in Go?

// import (
// 	"fmt"
// )

// func main() {
// 	var num interface{} = "42"
// 	intNum, ok := num.(int)

// 	fmt.Println("Potential Risks of Type Conversion in Go:")
// 	if ok {
// 		fmt.Printf("Converted integer: %d\n", intNum)
// 	} else {
// 		fmt.Println("Invalid type conversion")
// 	}
// }
// Options:

// a. Type conversion in Go is always safe and has no potential risks.
// b. Potential risks of type conversion include data loss and runtime panics if the conversion is not valid.
// c. Type conversion is a compile-time operation, so there are no runtime risks.
// d. Potential risks of type conversion include improved performance and reduced memory usage.
// Answer: b. Potential risks of type conversion include data loss and runtime panics if the conversion is not valid.

// Question 7: How can you convert a custom type to another type in Go?

// import "fmt"

// type Celsius float64
// type Fahrenheit float64

// func celsiusToFahrenheit(c Celsius) Fahrenheit {
// 	return Fahrenheit(c*9/5 + 32)
// }

// func main() {
// 	tempC := Celsius(25.0)
// 	tempF := celsiusToFahrenheit(tempC)

// 	fmt.Println("Converting Custom Types in Go:")
// 	fmt.Printf("%.2f°C is equal to %.2f°F\n", tempC, tempF)
// }
// Options:

// a. Custom type conversion in Go is not supported.
// b. You can only convert built-in types, not custom types.
// c. Custom type conversion is done by defining conversion functions between the custom types.
// d. Custom types are automatically converted to other types based on their underlying representation.
// Answer: c. Custom type conversion is done by defining conversion functions between the custom types.

// Question 8: Explain the difference between explicit and implicit type conversion.

// import (
// 	"fmt"
// )

// func main() {
// 	var a int = 42
// 	var b float64 = 3.14

// 	// Explicit type conversion
// 	result := float64(a) * b

// 	fmt.Println("Explicit vs. Implicit Type Conversion in Go:")
// 	fmt.Printf("Result (explicit): %.2f\n", result)

// 	// Implicit type conversion
// 	result2 := a * int(b)

// 	fmt.Printf("Result (implicit): %d\n", result2)
// }
// Options:

// a. Explicit type conversion is not allowed in Go, and implicit type conversion is always used.
// b. Explicit type conversion requires specifying the type, while implicit type conversion is automatic.
// c. Explicit type conversion is automatic, while implicit type conversion requires manual specification of the type.
// d. There is no difference between explicit and implicit type conversion in Go.
// Answer: b. Explicit type conversion requires specifying the type, while implicit type conversion is automatic.

// Question 9: What is the role of the interface{} type in type conversion?

// import (
// 	"fmt"
// )

// func printValue(val interface{}) {
// 	fmt.Println("Value:", val)
// }

// func main() {
// 	num := 42
// 	str := "Hello"

// 	fmt.Println("Role of interface{} in Type Conversion in Go:")
// 	printValue(num)
// 	printValue(str)
// }
// Options:

// a. The interface{} type is used for type conversion between custom types.
// b. The interface{} type is used to store values of any data type and is often used when the exact type is unknown.
// c. The interface{} type is a reserved keyword in Go and cannot be used for type conversion.
// d. The interface{} type is used exclusively for type assertion and not for type conversion.
// Answer: b. The interface{} type is used to store values of any data type and is often used when the exact type is unknown.

// Question 10: Can you perform type conversion between slices and arrays in Go?

// import (
// 	"fmt"
// )

// func main() {
// 	array := [3]int{1, 2, 3}
// 	slice := []int(array[:])

// 	fmt.Println("Type Conversion Between Slices and Arrays in Go:")
// 	fmt.Println("Array:", array)
// 	fmt.Println("Slice:", slice)
// }
// Options:

// a. Type conversion between slices and arrays is not supported in Go.
// b. Type conversion between slices and arrays is automatic and does not require explicit syntax.
// c. You can use the slice keyword to convert an array to a slice.
// d. Type conversion between slices and arrays requires a custom conversion function.
// Answer: b. Type conversion between slices and arrays is automatic and does not require explicit syntax.
