package fibonacci

func FibIt(n int) int {
	var (
		res  = 0
		calc = []int{n}
	)

	for len(calc) != 0 {
		c := calc[0]
		calc = append(calc[:0], calc[1:]...)
		if !(c <= 1) {
			calc = append(calc, c-1, c-2)
		} else {
			res++
		}
	}

	return res
}

func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

/*
 * Please note: This part is generated using Claude. 
*/

func FibMultithreaded(n int) uint64 {
	// Handle base cases
	if n <= 1 {
		return uint64(n)
	}

	// Create a channel to receive results from goroutines
	results := make(chan uint64)

	// Calculate number of goroutines to use
	// We'll use a goroutine for every 10 numbers to avoid overhead of too many goroutines
	chunkSize := 10
	numGoroutines := (n / chunkSize) + 1
	if numGoroutines > 10 {
		numGoroutines = 10 // Cap at 10 goroutines to prevent excessive resource usage
	}

	// Calculate how many numbers each goroutine should process
	numbersPerGoroutine := n / numGoroutines
	if n%numGoroutines != 0 {
		numbersPerGoroutine++
	}

	// Launch goroutines to calculate partial results
	for i := 0; i < numGoroutines; i++ {
		start := i * numbersPerGoroutine
		end := start + numbersPerGoroutine
		if end > n {
			end = n
		}

		go calculateFibRange(start, end, results)
	}

	// Collect and combine results
	var finalResult uint64
	for i := 0; i < numGoroutines; i++ {
		result := <-results
		if result > finalResult {
			finalResult = result
		}
	}

	return finalResult
}

// calculateFibRange calculates Fibonacci numbers for a given range
// and sends the largest result through the channel
func calculateFibRange(start, end int, results chan<- uint64) {
	if start <= 1 {
		start = 2 // Start from 2 since we know Fib(0)=0 and Fib(1)=1
	}

	// Initialize first two numbers
	a, b := uint64(0), uint64(1)

	// Calculate up to the start of our range
	for i := 2; i < start; i++ {
		a, b = b, a+b
	}

	// Calculate numbers in our range
	var largest uint64
	for i := start; i <= end; i++ {
		a, b = b, a+b
		if b > largest {
			largest = b
		}
	}

	results <- largest
}

