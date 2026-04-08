

// ============================================================
// TOPIC: Interfaces - Implicit Implementation, Polymorphism
// ============================================================


// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create an `Animal` struct and make it implement Speaker.
//     Pass it to the `present()` function.

// package main 
// import(
// 	"fmt"
// )

// type Speaker interface{
// 	Speak()string
// }

// type Animal struct{
// 	Name string


// }

// func  (a Animal) Speak()string  {

// 	return fmt.Sprintf("%s says Hello!", a.Name)
	
// }

// func present(s Speaker)  {
// 	fmt.Println(s.Speak())
// }


// func main(){
// 	animal:= Animal{Name: "Dog"}
// 	present(animal)
// }


//
// Q2: Create a `Shape` interface with methods:
//       Area() float64
//       Perimeter() float64
//     Implement it for Circle and Rectangle structs.
//     Write a function `printShape(s Shape)` that prints both.

// package main 
// import (
// 	"fmt"
// )


// type Shape interface{
// 	Name() string
// 	Area() float64
// 	Perimeter() float64
// }


// // Circle

// type Circle struct{
// 	name string
// 	radius float64
// }

// func (c Circle) Name() string{
// 	return c.name
// }

// func (c Circle) Area()float64{
// 	return 3.14*c.radius*c.radius
// }


// func (c Circle) Perimeter()float64{
// 	return 2*3.14*c.radius
// }



// // Rectangle


// type Rectangle struct{
// 	name string
// 	height float64
// 	width float64
// }

// func (r Rectangle) Name() string{
// 	return r.name
// }

// func (r Rectangle) Area()float64{
// 	return r.height*r.width	
// }

// func (r Rectangle) Perimeter()float64{
// 	return 2*(r.height+r.width)
// }



// // Polymorphic function

// func printShape(s Shape){
// 	fmt.Printf("Shape: %s\n", s.Name())
// 	fmt.Printf("Area: %.2f\n", s.Area())
// 	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
// }

// func main(){
// 	c:=Circle{name: "Circle", radius: 5}
// 	r:=Rectangle{name: "Rectangle", height: 4, width: 6}


// 	printShape(c)
// 	printShape(r)


// }



//
// Q3: Create a `Stringer` interface with `String() string`.
//     (This is actually built into Go as fmt.Stringer!)
//     Implement it on Person so fmt.Println prints nicely.



// package main 
// import(
// 	"fmt"
// )


// type Person struct{
// 	Name string
// 	Age int
// }


// func (p Person) String() string{
// 	return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
// }


// func main(){
// 	p:= Person{Name: "Shariful Islam", Age: 30}
// 	fmt.Println(p)
// }




//
// Q4: What happens if Robot is missing the Speak() method?
//     Try commenting it out - you'll get a compile error when
//     passing Robot to present(). This is Go's type safety!

// package main
// import (
// 	"fmt"
// )

// type Speaker interface{
// 	Speak() string
// }

// type Robot struct{
// 	Model string
// }

//  cannot use robot (variable of struct type Robot) as 
// Speaker value in argument to present: Robot does not 
// implement Speaker (missing method Speak)


// func present(s Speaker)  {
// 	fmt.Println(s.Speak())
// }

// func main()  {
// 	robot := Robot{Model: "T-800"}
// 	present(robot)
// }





//
// Q5: The empty interface `interface{}` (or `any` in Go 1.18+)
//     accepts any type. Write a function:
//       func describe(i any) string
//     that uses a type switch to describe the value.

// package main 
// import (
// 	"fmt"
// )


// func describe( i any)string{
// 	switch v:=i.(type) {
// 	case int:
// 		return fmt.Sprintf("Integer with the value %d", v)
// 	case float64:
// 		return fmt.Sprintf("Float with the value %f", v)
// 	case string:
// 		return fmt.Sprintf("String with the value %s", v)
// 	case bool:
// 		return fmt.Sprintf("Boolean with the value %t", v)
// 	case nil:
// 		return "nil value"
// 	default:
// 		return fmt.Sprintf("Invalid %T with the value %v", v, v)
		
// 	}
// }


// func main()  {

// 	fmt.Println(describe(120))
// 	fmt.Println(describe(10.25))
// 	fmt.Println(describe("I'm sharif who is web developer"))
// 	fmt.Println(describe(true))
// 	fmt.Println(describe(nil))
// 	fmt.Println(describe([]int {2,3,5,7,11}))
	

// }

//
// Q6: What is the difference between these two?
//       var s Speaker        // nil interface
//       var s Speaker = user // interface holding a User
//     (Hint: nil interface vs nil concrete value)
//
// var s Speaker        // nil interface: s is nil, no type, no value or type nil and value nil
// var s Speaker = user // interface holding a User: s is not nil, it holds a User value or type *User and value of user is nil

// fmt.Println(s == nil) // true for nil interface, false for interface holding a User
// ============================================================
