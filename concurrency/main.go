// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"time"
// 	// "sync"
// )

// ============================================================
// TOPIC: Concurrency - Goroutines, Channels, Context
// ============================================================

// --- EXAMPLE ---

// func main() {
// 	// --- GOROUTINES + CHANNELS ---
// 	// Channels let goroutines communicate safely
// 	results := make(chan string, 3) // buffered channel

// 	go fetchService("payments", results)
// 	go fetchService("analytics", results)
// 	go fetchService("notifications", results)

// 	// Collect results
// 	for i := 0; i < 3; i++ {
// 		fmt.Println("Result:", <-results)
// 	}

// 	// --- CONTEXT WITH TIMEOUT ---
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	ch := make(chan string, 1)
// 	go func() {
// 		time.Sleep(500 * time.Millisecond)
// 		ch <- "done"
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Got:", msg)
// 	case <-ctx.Done():
// 		fmt.Println("Timeout:", ctx.Err())
// 	}

// 	// --- PIPELINE PATTERN ---
// 	pipeline := make(chan int)
// 	squares := make(chan int)

// 	go producer(pipeline, 5)
// 	go squareWorker(pipeline, squares)

// 	// Consumer runs in main goroutine
// 	for v := range squares {
// 		fmt.Println("Square:", v)
// 	}
// }

// func fetchService(name string, out chan<- string) {
// 	delay := time.Duration(100+rand.Intn(300)) * time.Millisecond
// 	time.Sleep(delay)
// 	out <- fmt.Sprintf("%s responded in %s", name, delay)
// }

// func producer(out chan<- int, count int) {
// 	defer close(out)
// 	for i := 1; i <= count; i++ {
// 		out <- i
// 	}
// }

// func squareWorker(in <-chan int, out chan<- int) {
// 	defer close(out)
// 	for n := range in {
// 		out <- n * n
// 	}
// }

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Launch 5 goroutines that each print their index.
//     Use sync.WaitGroup to wait for all to finish.
//  var wg sync.WaitGroup;
 
// package main 
// import (
// 	"fmt"
// )

// func main(){
 
 
//  for i:=1;i<=5;i++{
//      wg.Add(1)
     
//      go func(i int){
//         defer wg.Done()
        
//         fmt.Println("Goroutine: ", i)
        
//      }(i)
//  }
// }
//  wg.Wait()
//  fmt.Println("All Goroutines are finished ")

//
// Q2: Create a channel of ints. Send numbers 1-10 in a
//     goroutine, close the channel, then range over it
//     in main to print them.

// package main
// import (
//     "fmt"
// )
// func main() {
//   ch:=make(chan int)
    
//     go func(){
//       for i:=1;i<=10;i++{
//           ch<-i
//       }  
//       close(ch)
//     }()
    
//     for num:=range ch{
//         fmt.Println(num)
//     }
//}
// Q3: Write a fan-out pattern: one producer sends work to
//     3 worker goroutines via a shared channel.




// package main
// import (
//     "fmt"
//     "sync"
  
//     )

// func main() {
    
//     jobs:=make(chan int)
//     var wg sync.WaitGroup
    
//     for w:=1;w<=3;w++{
//         wg.Add(1)
//         go worker(w, jobs, &wg)
//     }
    
//     go func(){
//       defer wg.Done()
//       for i:=1;i<=10;i++{
//           jobs<-i
//       }
       
//        close(jobs) 
//     }()
    
    
    
//     wg.Wait()
  
  
// }


// func worker(id int, jobs <-chan int,wg * sync.WaitGroup){
    
//     defer wg.Done()
    
//     for job :=range jobs{
//         fmt.Printf("Worker %d is processing job %d \n", id, job)
//     }
    
// }
//
// Q4: Use `select` with two channels and a timeout:
//       ch1 sends after 1s, ch2 sends after 2s,
//       timeout after 1.5s. What gets printed?

// package main 
// import (
// 	"fmt"
// 	"time"
// )


// func  main()  {
// 	cha1:=make(chan string)
// 	cha2:=make(chan string)

// 	go func(){
// 		time.Sleep(1*time.Second)
// 		cha1 <- "Message from channel 1"
// 	}()

// 	go func(){
// 		time.Sleep(2*time.Second)
// 		cha2 <- "Message from channel 2"
// 	}()


// 	for i:=0;i<3;i++{
// 	select{
// 	case msg:= <-cha1:
// 		fmt.Println(msg)
// 	case msg:= <-cha2:
// 		fmt.Println(msg)

// 	case <- time.After(1500*time.Millisecond):
// 		fmt.Println("timeout")

// 	}
// }

// }

//
// Q5: What happens if you send to an unbuffered channel
//     with no receiver? (Answer: deadlock!)
//     Try: ch := make(chan int); ch <- 1


// package main 
// import (
// 	// "fmt"
// )

// func main(){
// 	ch := make(chan int);

// 	ch <- 1

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:

// in unbuffered channels, need a receiver and a sender to proceed. 
// if there is no receiver, the sender will block indefinitely, leading to a deadlock.

// }
//
// Q6: Use context.WithCancel() to cancel a long-running
//     goroutine. The goroutine should check ctx.Done()
//     in a select loop.

package main 
import (
	"context"
	"fmt"
	"time"
)


func main(){
	ctx, cancel:=context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(2*time.Second)
	cancel()

	time.Sleep(1*time.Second)


}


func worker(ctx context.Context){
	for{
		select{
			case <-ctx.Done():
				fmt.Println("Worker received cancellation signal. Exiting.")
				return
			default:
				fmt.Println("Worker is doing work...")
				time.Sleep(500*time.Millisecond)
		}
	}
}
//
// ============================================================