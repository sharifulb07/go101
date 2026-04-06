package main

import (
	"fmt"

)

// ============================================================
// TOPIC: Go Basics - Variables, Types, Control Flow, Strings
// ============================================================

// --- EXAMPLE ---
func main (){

	//  Q1: Declare a constant called `maxRetries` with value 5.
	//     Print it inside main().
	const maxRetries int = 5
	fmt.Println("Value of maxRetries Constant :", maxRetries)

	// Q2: Write a function `isEven(n int) bool` that returns true
	//     if n is even. Call it from main() with a few test values.

	fmt.Println("Is 4 even?", isEven(4))
	fmt.Println("Is 7 even?", isEven(7))
	fmt.Println("Is 0 even?", isEven(0))
	fmt.Println("Is -2 even?", isEven(-2))
	fmt.Println("Is -3 even?", isEven(-3))

// Q3: Write a function `greet(name string) string` that returns
//     "Hello, <name>! Welcome to Go 101." using fmt.Sprintf.
	fmt.Println(greet("Hunaif Sharif"))
	fmt.Println(greet("Shariful Islam"))
	fmt.Println(greet("Ashia Khatun"))


	 //Q4: Use a for loop to print numbers 1 to 10. Skip even
	//     numbers using `continue`.

	for i:=1; i<=10;i++{
		if i%2==0{
			continue
		}
		fmt.Println("Number: ", i)
	}


	// Q5: Write a switch statement that takes a day string
//     ("Monday", "Saturday", etc.) and prints whether it's
//     a weekday or weekend.

fmt.Println(checkDay("Saturday"))
fmt.Println(checkDay("Monday"))
fmt.Println(checkDay("Birthday"))

// Q6: What will this print? Why?
//     for i := 0; i < 3; i++ {
//         defer fmt.Println(i)
//     }
//     (Answer: 2, 1, 0 - defer is LIFO)
//

 for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	// LIFO - Last In First Out
	// The deferred calls will be executed in reverse order of their appearance.
	//  if it starts from reverse order then it will print 2, 1, 0 that means the last value of i 2 less than 3 will be printed first and then 1 and then 0.


}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
//
//
// Q2: Write a function `isEven(n int) bool` that returns true
//     if n is even. Call it from main() with a few test values.


func isEven(n int) bool{
	if n % 2 == 0 {
		return true
	}
	return false
}
//
// Q3: Write a function `greet(name string) string` that returns
//     "Hello, <name>! Welcome to Go 101." using fmt.Sprintf.

func  greet(name string)string  {
	
	return fmt.Sprintf("Hello, %s! Welcome to Go 101.",name)
}
//
// Q4: Use a for loop to print numbers 1 to 10. Skip even
//     numbers using `continue`.



	// for i:=1; i<=10;i++{
	// 	if i%2==0{
	// 		continue
	// 	}
	// 	fmt.Println("Number: ", i)
	// }

//
// Q5: Write a switch statement that takes a day string
//     ("Monday", "Saturday", etc.) and prints whether it's
//     a weekday or weekend.

func checkDay(day string) string  {
	
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return fmt.Sprintf("%s is a Weekday", day )
	case "Saturday", "Sunday":
		return fmt.Sprintf("%s is a Weekend", day )
		default:
			return fmt.Sprintf("%s is an Invalid Day", day)
		
	}
}

//
// Q6: What will this print? Why?
//     for i := 0; i < 3; i++ {
//         defer fmt.Println(i)
//     }
//     (Answer: 2, 1, 0 - defer is LIFO)
//


	// LIFO - Last In First Out
	// The deferred calls will be executed in reverse order of their appearance.
	//  if it starts from reverse order then it will print 2, 1, 0 that means the last value of i 2 less than 3 will be printed first and then 1 and then 0.
// ============================================================
