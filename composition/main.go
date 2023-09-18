package main

// Composition:

// Question 1: What is composition in Go, and how is it different from inheritance?

// import "fmt"

// type Animal struct {
// 	name string
// }

// type Dog struct {
// 	Animal
// 	breed string
// }

// func main() {
// 	dog := Dog{
// 		Animal: Animal{name: "Buddy"},
// 		breed:  "Golden Retriever",
// 	}

// 	fmt.Println("Composition in Go:")
// 	fmt.Printf("Dog's name: %s\n", dog.name)
// 	fmt.Printf("Dog's breed: %s\n", dog.breed)
// }
// Options:

// a. Composition is a way to inherit behavior from parent classes, similar to traditional inheritance in other languages.
// b. Composition is a way to create new data types in Go, while inheritance is not supported.
// c. Composition is a way to combine multiple types to create a new type, while inheritance is a mechanism for code reuse and polymorphism.
// d. Composition is a synonym for inheritance in Go.
// Answer: c. Composition is a way to combine multiple types to create a new type, while inheritance is a mechanism for code reuse and polymorphism.

// Question 2: Explain how you can achieve composition in Go using structs.

// import "fmt"

// type Engine struct {
// 	Power int
// }

// type Car struct {
// 	Model  string
// 	Engine // Embedded struct
// }

// func main() {
// 	myCar := Car{
// 		Model: "Sedan",
// 		Engine: Engine{
// 			Power: 200,
// 		},
// 	}

// 	fmt.Println("Achieving Composition in Go:")
// 	fmt.Printf("Car Model: %s\n", myCar.Model)
// 	fmt.Printf("Engine Power: %d HP\n", myCar.Power)
// }
// Options:

// a. Composition in Go is achieved by directly inheriting from other structs.
// b. Composition in Go is achieved by using the extends keyword.
// c. Composition in Go is achieved by embedding one struct into another struct.
// d. Composition in Go is not supported.
// Answer: c. Composition in Go is achieved by embedding one struct into another struct.

// Question 3: How can you embed one struct into another to achieve composition?

// import "fmt"

// type Address struct {
// 	City  string
// 	State string
// }

// type Person struct {
// 	Name    string
// 	Address // Embedded struct
// }

// func main() {
// 	p := Person{
// 		Name: "Alice",
// 		Address: Address{
// 			City:  "New York",
// 			State: "NY",
// 		},
// 	}

// 	fmt.Println("Embedding Structs for Composition in Go:")
// 	fmt.Printf("Person: %s\n", p.Name)
// 	fmt.Printf("Address: %s, %s\n", p.City, p.State)
// }
// Options:

// a. Composition is achieved by creating separate structs with no relationship between them.
// b. Composition is achieved by using inheritance instead of embedding.
// c. Composition is achieved by embedding one struct type directly into another.
// d. Composition is achieved by using arrays instead of structs.
// Answer: c. Composition is achieved by embedding one struct type directly into another.

// Question 4: What is method overriding, and how does it work in composition?

// import "fmt"

// type Animal struct {
// 	Name string
// }

// func (a *Animal) Speak() {
// 	fmt.Println("Animal speaks")
// }

// type Dog struct {
// 	Animal
// 	Breed string
// }

// func (d *Dog) Speak() {
// 	fmt.Println("Dog barks")
// }

// func main() {
// 	dog := Dog{
// 		Animal: Animal{Name: "Buddy"},
// 		Breed:  "Golden Retriever",
// 	}

// 	fmt.Println("Method Overriding in Composition in Go:")
// 	dog.Speak()
// }
// Options:

// a. Method overriding is not supported in Go, and all methods have unique names.
// b. Method overriding allows a subclass to provide a specific implementation of a method that is already defined in its superclass.
// c. Method overriding is automatic in Go, and you don't need to specify it explicitly.
// d. Method overriding is achieved by using a separate interface for each subclass.
// Answer: b. Method overriding allows a subclass to provide a specific implementation of a method that is already defined in its superclass.

// Question 5: Provide an example of using interfaces for composition in Go.

// import "fmt"

// type Writer interface {
// 	Write(data string)
// }

// type ConsoleWriter struct{}

// func (cw ConsoleWriter) Write(data string) {
// 	fmt.Println("Writing to console:", data)
// }

// func main() {
// 	writer := ConsoleWriter{}
// 	writer.Write("Hello, Go!")

// 	fmt.Println("Using Interfaces for Composition in Go:")
// }
// Options:

// a. Composition with interfaces is not allowed in Go.
// b. Interfaces can only be used for inheritance, not composition.
// c. Interfaces can be used for composition by embedding them in structs.
// d. Interfaces are only used for defining method signatures, not for composition.
// Answer: c. Interfaces can be used for composition by embedding them in structs.

// Question 6: How does composition promote code reusability in Go?

// import "fmt"

// type DatabaseConnection struct {
// 	Host     string
// 	Port     int
// 	Username string
// 	Password string
// }

// func (db DatabaseConnection) Connect() {
// 	fmt.Printf("Connecting to %s:%d...\n", db.Host, db.Port)
// }

// type UserRepository struct {
// 	DB DatabaseConnection
// }

// func main() {
// 	db := DatabaseConnection{
// 		Host:     "localhost",
// 		Port:     5432,
// 		Username: "user",
// 		Password: "password",
// 	}

// 	userRepo := UserRepository{DB: db}
// 	userRepo.DB.Connect()

// 	fmt.Println("Promoting Code Reusability with Composition in Go:")
// }
// Options:

// a. Composition does not promote code reusability in Go.
// b. Code reusability is achieved only through inheritance, not composition.
// c. Composition allows you to reuse existing structs and their behavior within new structs, reducing code duplication.
// d. Code reusability in Go is solely dependent on the use of interfaces.
// Answer: c. Composition allows you to reuse existing structs and their behavior within new structs, reducing code duplication.

// Question 7: Can you compose multiple structs to create complex data structures?

// import "fmt"

// type Address struct {
// 	Street  string
// 	City    string
// 	State   string
// 	ZipCode string
// }

// type Person struct {
// 	Name    string
// 	Age     int
// 	Address Address
// }

// func main() {
// 	addr := Address{
// 		Street:  "123 Main St",
// 		City:    "New York",
// 		State:   "NY",
// 		ZipCode: "10001",
// 	}

// 	person := Person{
// 		Name:    "Alice",
// 		Age:     30,
// 		Address: addr,
// 	}

// 	fmt.Println("Composing Multiple Structs in Go:")
// 	fmt.Printf("Person: %s, Age: %d\n", person.Name, person.Age)
// 	fmt.Printf("Address: %s, %s, %s %s\n", person.Address.Street, person.Address.City, person.Address.State, person.Address.ZipCode)
// }
// Options:

// a. Composing multiple structs is not allowed in Go.
// b. Go does not support complex data structures created through composition.
// c. Yes, you can compose multiple structs to create complex data structures that represent hierarchical information.
// d. Complex data structures in Go are created exclusively using arrays and slices.
// Answer: c. Yes, you can compose multiple structs to create complex data structures that represent hierarchical information.
