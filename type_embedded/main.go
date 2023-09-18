package main

// Type Embedding:

// Question 1: What is type embedding in Go, and why is it useful?

// import "fmt"

// type Engine struct {
// 	Power int
// }

// type Car struct {
// 	Model string
// 	Engine
// }

// func main() {
// 	car := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	fmt.Println("Type Embedding in Go:")
// 	fmt.Printf("Car Model: %s\n", car.Model)
// 	fmt.Printf("Engine Power: %d HP\n", car.Power)
// }
// Options:

// a. Type embedding is a way to create new types by combining existing types, making it easier to manage complex data structures.
// b. Type embedding is a way to hide the fields and methods of one type within another type.
// c. Type embedding is a way to convert between different types in Go.
// d. Type embedding is not supported in Go.
// Answer: a. Type embedding is a way to create new types by combining existing types, making it easier to manage complex data structures.

// Question 2: How do you create an embedded type within a struct?

// import "fmt"

// type Engine struct {
// 	Power int
// }

// type Car struct {
// 	Model string
// 	Engine // Embedded type
// }

// func main() {
// 	myCar := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	fmt.Println("Creating an Embedded Type in Go:")
// 	fmt.Printf("Car Model: %s\n", myCar.Model)
// 	fmt.Printf("Engine Power: %d HP\n", myCar.Power)
// }
// Options:

// a. Embedded types are created by defining a new struct with fields from other structs.
// b. Embedded types are created using the embed keyword.
// c. Embedded types are created by simply declaring a field with the type you want to embed.
// d. Embedded types are created using inheritance.
// Answer: c. Embedded types are created by simply declaring a field with the type you want to embed.

// Question 3: Explain the concept of method promotion through type embedding.

// import "fmt"

// type Engine struct {
// 	Power int
// }

// func (e *Engine) Start() {
// 	fmt.Println("Engine started")
// }

// type Car struct {
// 	Model string
// 	Engine // Embedded type
// }

// func main() {
// 	myCar := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	fmt.Println("Method Promotion Through Type Embedding in Go:")
// 	myCar.Start()
// }
// Options:

// a. Method promotion allows methods of an embedded type to be called directly on the outer type, as if they were declared on the outer type.
// b. Method promotion is a technique to hide methods of an embedded type in the outer type.
// c. Method promotion is a way to call methods of an embedded type using a special syntax.
// d. Method promotion is not supported in Go.
// Answer: a. Method promotion allows methods of an embedded type to be called directly on the outer type, as if they were declared on the outer type.

// Question 4: What happens when multiple embedded types have methods with the same name?

// import "fmt"

// type Engine struct {
// 	Power int
// }

// func (e *Engine) Start() {
// 	fmt.Println("Engine started")
// }

// type Radio struct{}

// func (r *Radio) Start() {
// 	fmt.Println("Radio started")
// }

// type Car struct {
// 	Model string
// 	Engine
// 	Radio
// }

// func main() {
// 	myCar := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 		Radio: Radio{},
// 	}

// 	fmt.Println("Handling Multiple Embedded Types with Same Method Names in Go:")
// 	myCar.Engine.Start()
// 	myCar.Radio.Start()
// }
// Options:

// a. It is not possible to have multiple embedded types with methods of the same name in Go.
// b. The method of the last embedded type with the same name takes precedence and shadows the methods of the previous types.
// c. All methods with the same name from embedded types are available, and you can choose which one to call explicitly.
// d. The methods of embedded types with the same name are combined into a single method in the outer type.
// Answer: b. The method of the last embedded type with the same name takes precedence and shadows the methods of the previous types.

// Question 5: How can you access the fields and methods of an embedded type?

// import "fmt"

// type Engine struct {
// 	Power int
// }

// func (e *Engine) Start() {
// 	fmt.Println("Engine started")
// }

// type Car struct {
// 	Model string
// 	Engine
// }

// func main() {
// 	myCar := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	fmt.Println("Accessing Fields and Methods of an Embedded Type in Go:")
// 	fmt.Printf("Car Model: %s\n", myCar.Model)
// 	fmt.Printf("Engine Power: %d HP\n", myCar.Power)
// 	myCar.Start()
// }
// Options:

// a. Fields and methods of an embedded type cannot be accessed directly.
// b. Fields and methods of an embedded type are accessed using the dot notation on the outer type.
// c. Fields of an embedded type are accessible, but methods are not.
// d. Fields and methods of an embedded type are accessed using a special keyword.
// Answer: b. Fields and methods of an embedded type are accessed using the dot notation on the outer type.

// Question 6: What is the difference between embedded types and traditional composition?

// import "fmt"

// type Engine struct {
// 	Power int
// }

// type Car struct {
// 	Model string
// 	Engine
// }

// type Boat struct {
// 	Model string
// 	Engine Engine
// }

// func main() {
// 	car := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	boat := Boat{
// 		Model: "Yacht",
// 		Engine: Engine{
// 			Power: 500,
// 		},
// 	}

// 	fmt.Println("Difference Between Embedded Types and Traditional Composition in Go:")
// 	fmt.Printf("Car Engine Power: %d HP\n", car.Power)
// 	fmt.Printf("Boat Engine Power: %d HP\n", boat.Engine.Power)
// }
// Options:

// a. There is no difference between embedded types and traditional composition; they are used interchangeably.
// b. Embedded types allow you to access fields and methods of the embedded type directly, while traditional composition requires explicit delegation.
// c. Traditional composition is not supported in Go; only embedded types are used for composition.
// d. Traditional composition is a more efficient way to compose types compared to embedded types.
// Answer: b. Embedded types allow you to access fields and methods of the embedded type directly, while traditional composition requires explicit delegation.

// Question 7: Can you embed an interface within a struct in Go?

// import "fmt"

// type Writer interface {
// 	Write(data string)
// }

// type ConsoleWriter struct{}

// func (cw ConsoleWriter) Write(data string) {
// 	fmt.Println("Writing to console:", data)
// }

// type Logger struct {
// 	Writer // Embedded interface
// }

// func main() {
// 	logger := Logger{Writer: ConsoleWriter{}}
// 	logger.Write("Hello, Go!")

// 	fmt.Println("Embedding an Interface Within a Struct in Go:")
// }
// Options:

// a. Embedding an interface within a struct is not allowed in Go.
// b. Interfaces can only be embedded within other interfaces, not structs.
// c. Yes, you can embed an interface within a struct in Go.
// d. Interfaces should not be used with embedding; they should be explicitly implemented.
// Answer: c. Yes, you can embed an interface within a struct in Go.

// Question 8: How does type embedding relate to code organization and reuse?

// import "fmt"

// type Engine struct {
// 	Power int
// }

// type Car struct {
// 	Model string
// 	Engine // Embedded type
// }

// func main() {
// 	myCar := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	fmt.Println("Type Embedding, Code Organization, and Reuse in Go:")
// 	fmt.Printf("Car Model: %s\n", myCar.Model)
// 	fmt.Printf("Engine Power: %d HP\n", myCar.Power)
// }
// Options:

// a. Type embedding has no impact on code organization and reuse.
// b. Type embedding promotes better code organization by encapsulating related data and behavior within a single type.
// c. Code organization and reuse are only achieved through inheritance, not type embedding.
// d. Type embedding can only be used for code reuse, not for code organization.
// Answer: b. Type embedding promotes better code organization by encapsulating related data and behavior within a single type.

// Question 9: What are the advantages and disadvantages of using type embedding in Go?

// import "fmt"

// type Engine struct {
// 	Power int
// }

// type Car struct {
// 	Model string
// 	Engine // Embedded type
// }

// func main() {
// 	myCar := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	fmt.Println("Advantages and Disadvantages of Using Type Embedding in Go:")
// 	fmt.Printf("Car Model: %s\n", myCar.Model)
// 	fmt.Printf("Engine Power: %d HP\n", myCar.Power)
// }
// Options:

// a. The advantages of type embedding include code reuse, improved code organization, and better encapsulation. The disadvantages include increased complexity and reduced flexibility.
// b. The advantages of type embedding include simplicity and reduced complexity. The disadvantages include code duplication and lack of flexibility.
// c. Type embedding has no advantages or disadvantages; it is a neutral feature in Go.
// d. The advantages of type embedding include better performance and reduced memory usage. The disadvantages include increased code verbosity.
// Answer: a. The advantages of type embedding include code reuse, improved code organization, and better encapsulation. The disadvantages include increased complexity and reduced flexibility.

// Question 10: Provide an example of a real-world use case for type embedding in Go.

// import "fmt"

// type Employee struct {
// 	ID   int
// 	Name string
// }

// type Manager struct {
// 	Employee // Embedded type
// 	TeamSize int
// }

// func main() {
// 	manager := Manager{
// 		Employee: Employee{
// 			ID:   101,
// 			Name: "Alice",
// 		},
// 		TeamSize: 5,
// 	}

// 	fmt.Println("Real-World Use Case for Type Embedding in Go:")
// 	fmt.Printf("Manager ID: %d\n", manager.ID)
// 	fmt.Printf("Manager Name: %s\n", manager.Name)
// 	fmt.Printf("Team Size: %d\n", manager.TeamSize)
// }
// Options:

// a. Type embedding is not applicable in real-world scenarios.
// b. A real-world use case for type embedding is modeling hierarchical relationships, such as employees and managers.
// c. Type embedding is only used in academic exercises, not in practical programming.
// d. Type embedding is primarily used for low-level system programming.
// Answer: b. A real-world use case for type embedding is modeling hierarchical relationships, such as employees and managers.
