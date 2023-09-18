package main

// 1. What is type conversion in Go, and why is it necessary?

// package main

// import "fmt"

// func main() {
//     var x int = 42
//     var y float64 = float64(x) // Type conversion from int to float64
//     fmt.Println("Converted value:", y)
// }

// 2. How do you convert between different numeric types in Go?

// package main

// import "fmt"

// func main() {
//     var a int = 42
//     var b float64 = float64(a) // Conversion from int to float64
//     fmt.Println("Converted value:", b)
// }

// 3.  Explain type conversion between strings and other data types.

// package main

// import (
//     "fmt"
//     "strconv"
// )

// func main() {
//     str := "42"

//     // String to integer
//     num, _ := strconv.Atoi(str)
//     fmt.Println("String to Int:", num)

//     // Integer to string
//     str2 := strconv.Itoa(42)
//     fmt.Println("Int to String:", str2)
// }

// 4. What is type assertion, and when is it used?

// Type assertion is used to check the actual type of an interface value and extract the underlying value of a specific type.
// It is used when working with interfaces that may hold values of different types.

// package main

// import "fmt"

// func main() {
//     var val interface{} = 42
//     x, ok := val.(int) // Type assertion to check if "val" is an int
//     if ok {
//         fmt.Println("Type assertion result:", x)
//     } else {
//         fmt.Println("Type assertion failed.")
//     }
// }

// 5. How can you check if a type conversion is valid in Go?

// package main

// import "fmt"

// func main() {
//     var val interface{} = 42
//     x, ok := val.(string) // Type assertion to check if "val" is a string
//     if ok {
//         fmt.Println("Type assertion result:", x)
//     } else {
//         fmt.Println("Type assertion failed.")
//     }
// }

// 6. What are the potential risks of type conversion in Go?

// 7. How can you convert a custom type to another type in Go?

// Custom types can be converted to other types by defining conversion methods for them.

// package main

// import "fmt"

// type Celsius float64
// type Fahrenheit float64

// func cToF(c Celsius) Fahrenheit {
//     return Fahrenheit(c*9/5 + 32)
// }

// func main() {
//     temperatureC := Celsius(25.5)
//     temperatureF := cToF(temperatureC)
//     fmt.Printf("%.2f°C is equal to %.2f°F\n", temperatureC, temperatureF)
// }

// 8. Explain the difference between explicit and implicit type conversion.

// 9. What is the role of the interface{} type in type conversion?

// 10. Can you perform type conversion between slices and arrays in Go?
