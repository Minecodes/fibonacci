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
