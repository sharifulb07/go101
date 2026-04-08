// package main

// import "fmt"

// // ============================================================
// // TOPIC: Functions - Multiple Returns, Variadic, Closures
// // ============================================================

// // --- EXAMPLE ---

// func main() {
// 	// Basic function call
// 	sum := add(10, 32)
// 	fmt.Println("sum:", sum)

// 	// Multiple return values + error handling
// 	quotient, remainder, err := divide(100, 7)
// 	if err != nil {
// 		fmt.Println("divide error:", err)
// 		return
// 	}
// 	fmt.Printf("100 / 7 = %d remainder %d\n", quotient, remainder)

// 	// Variadic function
// 	total := variadicSum(1, 2, 3, 4, 5)
// 	fmt.Println("variadic total:", total)

// 	// Closures - function that captures outer variable
// 	counter := makeCounter()
// 	fmt.Println(counter()) // 1
// 	fmt.Println(counter()) // 2
// 	fmt.Println(counter()) // 3

// 	// Anonymous function
// 	double := func(n int) int { return n * 2 }
// 	fmt.Println("double 5:", double(5))
// }

// func add(a, b int) int {
// 	return a + b
// }

// func divide(dividend, divisor int) (int, int, error) {
// 	if divisor == 0 {
// 		return 0, 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return dividend / divisor, dividend % divisor, nil
// }

// func variadicSum(values ...int) int {
// 	total := 0
// 	for _, v := range values {
// 		total += v
// 	}
// 	return total
// }

// func makeCounter() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }




// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Write a function `multiply(a, b int) int` and call it.
// package main
// import (
//     "fmt"
//     )
    
    
// func main(){
//     fmt.Println(multiply(301,48))
// }

// func multiply(a, b int)int{
//     return a*b
// }

//
// Q2: Write a function `swap(a, b string) (string, string)`
//     that returns the two strings in reverse order.

// package main
// import (
//     "fmt"
//     )
    
    
// func main(){
//     fmt.Println(swap("sharif", "adiyat"))
// }

// func swap(a, b string)(string, string){
//     return b,a
// }


//
// Q3: Write a function `max(nums ...int) int` that returns
//     the largest number from a variadic list.

// package main
// import (
//     "fmt"
//     )
    
    
// func main(){
//     fmt.Println(max(50,8,6,3,4,100,600,700,10,40))
// }


// func max(nums ...int)int{
//     largest:=0
//     for _,num:=range nums{
//         if num>largest{
//             largest=num
//         }
//     }
    
//     return largest
// }



//
// Q4: Write a closure `makeAdder(x int) func(int) int` that
//     returns a function which adds x to its argument.
//     Example: addFive := makeAdder(5); addFive(3) -> 8


// package main
// import (
//     "fmt"
//     )
    
    
// func main(){
//     addFive:=makeAdder(5)
//     addTen:=makeAdder(10)
    
   
//    fmt.Println(addFive(45))
//    fmt.Println(addTen(78))
   
// }


// func makeAdder(x int) func(int)int{
    
//     return func(y int)int{
//         return x+y
//     }
// }


//
// Q5: Write a function `apply(a, b int, op func(int,int) int) int`
//     that takes two ints and a function, and applies the
//     function. Pass in add, multiply, etc.
//

// package main
// import (
//     "fmt"
//     )
    
    
// func main(){
    
    
//  fmt.Println("Add:       ", apply(20, 30, add))        // 50
// 	fmt.Println("Subtract:  ", apply(50, 30, subtraction)) // 20
// 	fmt.Println("Multiply:  ", apply(30, 60, multiply))   // 1800
// 	fmt.Println("Divide:    ", apply(3000, 40, divide)) 
   
   
// }



// // higher order function concept 

// func apply(a,b int, op func(int, int)int)int{
//     return op(a,b)
// }

// func add(a,b int)int{
//     return a+b
// }

// func multiply(a,b int)int{
//     return a*b
// }

// func subtraction(a,b int)int{
//     return a-b
// }


// func divide(a, b int)int {
//     if b==0{
//         panic("A number cannot be divided by 0")
//     }
//     return a/b
    
// }






// Q6: Rewrite `divide` using named return values:
//     func divide(a, b int) (q int, r int, err error)
//

package main
import (
	"fmt"
	"errors"
)


func divide(a, b int ) (q int, r int, err error){

	if b==0{
		err:=errors.New("cann't divide by zero")
		return 0, 0, err
	}else{
		q=a/b
		r=a%b
		return q, r, nil
	}
}


func main(){
	q, r, err := divide(100, 7)
	if err != nil{
		fmt.Println("divide error:", err)
		return
	}
	fmt.Printf("100 / 7 = %d remainder %d\n", q, r)


	q, r, err = divide(100, 0)
	if err != nil{
		fmt.Println("divide error:", err)
		return
	}
	fmt.Printf("100 / 0 = %d remainder %d\n", q, r)
}



// ============================================================
