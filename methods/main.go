// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // ============================================================
// // TOPIC: Methods - Value vs Pointer Receivers
// // ============================================================

// // --- EXAMPLE ---

// type Notebook struct {
// 	Title  string
// 	Notes  []string
// 	Closed bool
// }

// // Value receiver - does NOT modify the original
// func (n Notebook) Count() int {
// 	return len(n.Notes)
// }

// // Value receiver - reads but does not modify
// func (n Notebook) Summary() string {
// 	return fmt.Sprintf("%s: %s", n.Title, strings.Join(n.Notes, ", "))
// }

// // Pointer receiver - DOES modify the original
// func (n *Notebook) Add(note string) {
// 	if n.Closed {
// 		fmt.Println("notebook is closed; cannot add note")
// 		return
// 	}
// 	n.Notes = append(n.Notes, note)
// }

// // Pointer receiver - modifies state
// func (n *Notebook) Close() {
// 	n.Closed = true
// }

// func main() {
// 	nb := &Notebook{Title: "Go Webinar"}

// 	nb.Add("functions return multiple values")
// 	nb.Add("struct embedding promotes fields")

// 	fmt.Println("notes count:", nb.Count())
// 	fmt.Println(nb.Summary())

// 	nb.Close()
// 	nb.Add("this will not be added") // closed!
// }

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create a `Counter` struct with a `Value int` field.
//     Add methods:
//       - Increment() - pointer receiver, adds 1
//       - Reset()     - pointer receiver, sets to 0
//       - Get() int   - value receiver, returns Value

// package main
// import "fmt"

// type Counter struct {
// 	Value int
// }


// func (c *Counter) Increment() {
// 	c.Value++
// }

// func (c *Counter) Reset() {
// 	c.Value = 0
// }

// func (c Counter) Get() int {
// 	return c.Value
// }


// func main() {
	

//   c:=&Counter{Value:2}
//     c.Increment()
//     c.Increment()
//     c.Increment()
//     c.Increment()
//     c.Increment()
//     fmt.Println(c.Get())
//    c.Reset()
//    fmt.Println(c.Get())
//        c.Increment()
//     c.Increment()
//     c.Increment()
//     fmt.Println(c.Get())// Should print 2
// }

//
// Q2: Create a `Rectangle` struct with Width and Height.
//     Add methods:
//       - Area() float64
//       - Perimeter() float64
//       - Scale(factor float64) - pointer receiver

// package main 
// import (
// 	"fmt"
// )

// type Rectangle struct{
// 	Height float64
// 	Width float64
// }


// func (r Rectangle)Area()float64{
// 	return r.Height*r.Width
// }

// func (r Rectangle)Perimeter()float64{
// 	return 2*(r.Height+r.Width)
// }

// func (r *Rectangle)Scale(factor float64){
// 	r.Height*=factor
// 	r.Width*=factor
// }


// func main(){
// 	r:=&Rectangle{Height:2,Width:3}
// 	fmt.Println("Area:", r.Area())
// 	fmt.Println("Perimeter:", r.Perimeter())
// 	r.Scale(2)
// 	fmt.Println("After Scale Set")
// 	fmt.Println("Scaled Area:", r.Area())
// 	fmt.Println("Scaled Perimeter:", r.Perimeter())

// }



//
// Q3: Why does Add() use a pointer receiver (*Notebook) but
//     Count() uses a value receiver (Notebook)?
//     (Answer: Add modifies the struct, Count only reads it)
// package main

// import (
// 	"fmt"
// 	"strings"
// )



// type Notebook struct {
// 	Title  string
// 	Notes  []string
// 	Closed bool
// }

// // Value receiver - does NOT modify the original
// func (n Notebook) Count() int {
// 	// Count function returns the number of notes in the notebook without modifying the notebook itself, so it uses a value receiver.
// 	return len(n.Notes)
// }

// // Value receiver - reads but does not modify
// func (n Notebook) Summary() string {
// 	return fmt.Sprintf("%s: %s", n.Title, strings.Join(n.Notes, ", "))
// }

// // Pointer receiver - DOES modify the original
// func (n *Notebook) Add(note string) {
// 	if n.Closed {
// 		fmt.Println("notebook is closed; cannot add note")
// 		return
// 	}
// 	// Add function modifies the notebook by adding a new note to the Notes slice, so it uses a pointer receiver to ensure that the changes are reflected in the original notebook.
// 	// here new note is appended to the Notes slice of the notebook that means changes the Notes field of the notebook, so it needs to be a pointer receiver to modify the original notebook.
// 	n.Notes = append(n.Notes, note)
// }

// // Pointer receiver - modifies state
// func (n *Notebook) Close() {
	
// 	//n.Closed=true sets the Closed field of the notebook to true, which modifies the state of the notebook. Therefore, it uses a pointer receiver to ensure that the change is reflected in the original notebook.

// 	n.Closed = true
// }

// func main() {
// 	nb := &Notebook{Title: "Go Webinar"}

// 	nb.Add("functions return multiple values")
// 	nb.Add("struct embedding promotes fields")

// 	fmt.Println("notes count:", nb.Count())
// 	fmt.Println(nb.Summary())

// 	nb.Close()
// 	nb.Add("this will not be added") // closed!
// }

//
// Q4: What happens if you change Add() to a value receiver?
//     Try it and see - the notes won't actually be saved!


// package main

// import (
// 	"fmt"
// 	"strings"
// )



// type Notebook struct {
// 	Title  string
// 	Notes  []string
// 	Closed bool
// }

// // Value receiver - does NOT modify the original
// func (n Notebook) Count() int {
	
// 	return len(n.Notes)
// }

// // Value receiver - reads but does not modify
// func (n Notebook) Summary() string {
// 	return fmt.Sprintf("%s: %s", n.Title, strings.Join(n.Notes, ", "))
// }

// // In Add function value receiver - DOES not allow to modify the original
// // so no new note will be added to the Notes slice of the original notebook, and the changes will not be reflected in the original notebook. Therefore, if you change Add() to a value receiver, the notes won't actually be saved in the original notebook.
// func (n Notebook) Add(note string) {
// 	if n.Closed {
// 		fmt.Println("notebook is closed; cannot add note")
// 		return
// 	}
	
// 	n.Notes = append(n.Notes, note)
// }

// // Pointer receiver - modifies state
// func (n *Notebook) Close() {
	
	

// 	n.Closed = true
// }

// func main() {
// 	nb := &Notebook{Title: "Go Webinar"}

// 	nb.Add("functions return multiple values")
// 	nb.Add("struct embedding promotes fields")

// 	fmt.Println("notes count:", nb.Count())
// 	fmt.Println(nb.Summary())

// 	nb.Close()
// 	nb.Add("this will not be added") // closed!
// }


//
// Q5: Add a method `Remove(index int)` to Notebook that
//     removes a note by index (use pointer receiver).
//     Hint: use append(n.Notes[:i], n.Notes[i+1:]...)


package main

import (
	"fmt"
	"strings"
)



type Notebook struct {
	Title  string
	Notes  []string
	Closed bool
}

// Value receiver - does NOT modify the original
func (n Notebook) Count() int {
	
	return len(n.Notes)
}

// Value receiver - reads but does not modify
func (n Notebook) Summary() string {
	return fmt.Sprintf("%s: %s", n.Title, strings.Join(n.Notes, ", "))
}


func (n *Notebook) Add(note string) {
	if n.Closed {
		fmt.Println("notebook is closed; cannot add note")
		return
	}
	
	n.Notes = append(n.Notes, note)
}

func (n *Notebook) Remove(index int) {
	if index < 0 || index >= len(n.Notes) {
		fmt.Println("invalid index; cannot remove note")
		return
	}
	n.Notes = append(n.Notes[:index], n.Notes[index+1:]...)
}


// Pointer receiver - modifies state
func (n *Notebook) Close() {
	
	

	n.Closed = true
}

func main() {
	nb := &Notebook{Title: "Go Webinar"}

	nb.Add("functions return multiple values")
	nb.Add("struct embedding promotes fields")
	nb.Add("struct values are copied on method calls")
	nb.Add("struct pointers allow modification in methods")

	fmt.Println("notes count:", nb.Count())
	fmt.Println(nb.Summary())

	nb.Remove(2)
	fmt.Println("After removing index 2:")
	fmt.Println("notes count:", nb.Count())
	fmt.Println(nb.Summary())

	nb.Close()
	nb.Add("this will not be added") // closed!
}

//
// ============================================================
