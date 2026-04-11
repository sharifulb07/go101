// package main

// import "fmt"

// // ============================================================
// // TOPIC: Structs - Composition, Embedding, Nested Structs
// // ============================================================

// // --- EXAMPLE ---

// type Address struct {
// 	City    string
// 	Country string
// }

// type Person struct {
// 	Name    string
// 	Age     int
// 	Address Address // nested struct
// }

// type Employee struct {
// 	Person       // embedded struct (field promotion)
// 	Role   string
// 	Salary int
// }

// func main() {
// 	// Struct literal
// 	person := Person{
// 		Name: "Rob",
// 		Age:  45,
// 		Address: Address{
// 			City:    "San Francisco",
// 			Country: "USA",
// 		},
// 	}
// 	fmt.Printf("Person: %+v\n", person)
// 	fmt.Println("City:", person.Address.City)

// 	// Struct embedding - fields are promoted
// 	employee := Employee{
// 		Person: person,
// 		Role:   "Engineer",
// 		Salary: 160000,
// 	}

// 	// Access promoted fields directly
// 	fmt.Println("Name:", employee.Name)           // promoted from Person
// 	fmt.Println("City:", employee.Address.City)    // promoted through Person
// 	fmt.Printf("Full: %s, %s, $%d\n", employee.Name, employee.Role, employee.Salary)

// 	// Structs are value types (copied on assignment)
// 	p2 := person
// 	p2.Name = "Alice"
// 	fmt.Println("Original:", person.Name)  // still "Rob"
// 	fmt.Println("Copy:", p2.Name)          // "Alice"
// }

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create a `Book` struct with Title, Author, and Pages.
//     Create an instance and print it.


// package main 
// import (
// 	"fmt"
// )


// type Book struct{
// 	Title string
// 	Author string
// 	Pages int 
// }



// func main(){
// 	book1:=Book{
// 		Title: "First Attempt",
// 		Author: "Hunaif Sharif",
// 		Pages: 352,
// 	}

// 	fmt.Printf(" Title: %s\n Author: %s \n Pages:%d \n", book1.Title, book1.Author, book1.Pages)

// }


//
// Q2: Create a `Library` struct that has a Name and a slice
//     of Books ([]Book). Add 3 books and loop through them.

// package main 
// import (
// 	"fmt"
// )


// type Book struct{
// 	Title string
// 	Author string
// 	Pages int 
// }

// type Library struct{
// 	Name string
// 	Books []Book
// }


// func  main()  {
// 	l1:=Library{
// 		Name: "City Library",
// 		Books: []Book{
// 			{Title: "Basic Go", Author: "S M Shariar Islam", Pages: 360},
// 			{Title: "Advanced Go", Author: "Alice Bob", Pages: 200},
// 			{Title: "Practical go lang", Author: "Shariful Islam", Pages: 420},
// 		},
// 	}

// 	fmt.Println("The name of the Library: ", l1.Name)

// 	for i, book :=range l1.Books{
// 		fmt.Printf("Book %d \n", i)
// 		fmt.Println("Title: ", book.Title)
// 		fmt.Println("Author: ", book.Author)
// 		fmt.Println("Pages: ", book.Pages)
// 	}
	
// }

//
// Q3: Create a `Student` struct that embeds `Person` and adds
//     a `GPA` field. Show that you can access Name directly.

// package main 
// import (
// 	"fmt"
// )


// type Person struct{
// 	Name string 
// 	Age int 
// }


// type Student struct{
// 	Person
// 	GPA float64
// }



// func main(){
// 	s1:=Student{
// 		Person:Person {
// 			Name:"Shariful islam",
// 			Age: 35,
// 		},
// 		GPA: 3.50,
// 	}

// 	fmt.Println("Name: ", s1.Name)
// 	fmt.Println("Age: ", s1.Age)
// 	fmt.Println("GPA: ", s1.GPA)
// }


//
// Q4: Write a function `newPerson(name string, age int) *Person`
//     that returns a pointer to a new Person. Why might you
//     return a pointer instead of a value?


// package main 
// import (
// 	"fmt"
// )


// type Person struct{
// 	Name string
// 	Age int
// }



// // Return a value copies the entire Person Struct. For a large struct , it is inefficient.

// //a pointer passes just a memory address which is very cheap


// func newPerson(name string, age int) *Person{
// 	return &Person{
// 		Name: name,
// 		Age: age,
// 	}
// }





// func main ()  {
// 	p1:=newPerson("Hunaif Sharif", 07)

// 	fmt.Println(p1.Name)
// 	fmt.Println(p1.Age)
// }



//
// Q5: What happens when you compare two structs with ==?
//     Try comparing two Person values. Does it work?
//     (Hint: it works if all fields are comparable)


// package main 
// import (
// 	"fmt"
// )


// type Person struct{
// 	Name string
// 	Age int 
// }


// func main(){

// 	// always both value of two structs will be same

// // Struct comparison works only if all fields are comparable types, such as: int , float , string , bool, array, pointer,

// // When Struct does NOT work such as: slice, maps, functions 

// 	p1:=Person{Name: "sharif", Age: 32}
// 	p2:=Person{Name: "sharif", Age: 32}
// 	p3:=Person{Name: "hunaif", Age: 6}

// 	fmt.Println(p1==p2) // true

// 	fmt.Println(p2==p3) //false
// }




//
// Q6: Create an anonymous struct inline:
//     config := struct{ Host string; Port int }{"localhost", 8080}


package main 
import (
	"fmt"
)


func main ()  {
	
	// an anonymous struct
	config:=struct{ Host string; Port int}{"localhost", 8080}

	fmt.Println("Host: ", config.Host)
	fmt.Println("Port: ", config.Port)
}
//


// ============================================================
