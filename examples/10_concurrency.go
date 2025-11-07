package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Contoh Concurrency: Goroutine dan Channel ===\n")

	// Basic goroutine
	fmt.Println("=== Basic Goroutine ===")
	go sayHello("Goroutine 1")
	go sayHello("Goroutine 2")
	go sayHello("Goroutine 3")
	
	// Sleep untuk memberi waktu goroutine selesai
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// Anonymous goroutine
	fmt.Println("=== Anonymous Goroutine ===")
	go func() {
		fmt.Println("Anonymous goroutine running")
	}()
	time.Sleep(50 * time.Millisecond)
	fmt.Println()

	// Goroutine dengan parameter
	fmt.Println("=== Goroutine dengan Parameter ===")
	for i := 1; i <= 3; i++ {
		go printNumber(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// Channel
	fmt.Println("=== Basic Channel ===")
	ch := make(chan string)
	
	go func() {
		ch <- "Hello from channel"
	}()
	
	message := <-ch
	fmt.Println("Received:", message)
	fmt.Println()

	// Channel dengan goroutine
	fmt.Println("=== Channel Communication ===")
	messages := make(chan string)
	
	go sendMessages(messages)
	
	msg1 := <-messages
	msg2 := <-messages
	fmt.Println("Received:", msg1)
	fmt.Println("Received:", msg2)
	fmt.Println()

	// Buffered channel
	fmt.Println("=== Buffered Channel ===")
	buffered := make(chan int, 3)
	
	buffered <- 1
	buffered <- 2
	buffered <- 3
	// Tidak perlu goroutine karena channel di-buffer
	
	fmt.Println("Received:", <-buffered)
	fmt.Println("Received:", <-buffered)
	fmt.Println("Received:", <-buffered)
	fmt.Println()

	// Channel dengan range
	fmt.Println("=== Channel dengan Range ===")
	numbers := make(chan int)
	
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // Penting: close channel setelah selesai
	}()
	
	for num := range numbers {
		fmt.Printf("Received number: %d\n", num)
	}
	fmt.Println()

	// Select statement
	fmt.Println("=== Select Statement ===")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
	fmt.Println()

	// Select dengan default
	fmt.Println("=== Select dengan Default ===")
	ch3 := make(chan string)
	
	select {
	case msg := <-ch3:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received (non-blocking)")
	}
	fmt.Println()

	// WaitGroup
	fmt.Println("=== WaitGroup ===")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All workers completed")
	fmt.Println()

	// Worker pool pattern
	fmt.Println("=== Worker Pool ===")
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Println()

	// Mutex untuk sync
	fmt.Println("=== Mutex ===")
	demonstrateMutex()
	fmt.Println()

	// Channel direction
	fmt.Println("=== Channel Direction ===")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	
	ping(pings, "Hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println()

	// Timeout pattern
	fmt.Println("=== Timeout Pattern ===")
	ch4 := make(chan string, 1)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch4 <- "Result"
	}()
	
	select {
	case res := <-ch4:
		fmt.Println("Received:", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout!")
	}
	fmt.Println()

	// Ticker
	fmt.Println("=== Ticker ===")
	ticker := time.NewTicker(50 * time.Millisecond)
	done := make(chan bool)
	
	go func() {
		count := 0
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				count++
				fmt.Printf("Tick at %v (count: %d)\n", t.Format("15:04:05.000"), count)
			}
		}
	}()
	
	time.Sleep(200 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func sayHello(name string) {
	fmt.Printf("Hello from %s\n", name)
}

func printNumber(num int) {
	fmt.Printf("Number: %d\n", num)
}

func sendMessages(ch chan string) {
	ch <- "First message"
	ch <- "Second message"
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(50 * time.Millisecond)
		results <- j * 2
	}
}

// Mutex example
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func demonstrateMutex() {
	counter := SafeCounter{value: 0}
	var wg sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())
}

// Channel dengan direction (send-only dan receive-only)
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
