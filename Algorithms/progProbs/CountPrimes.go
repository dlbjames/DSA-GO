package main

import "fmt"

func main() {
	fmt.Println(primesUpTo(20))
}

func primesUpTo(n int) []int {
	var primes []bool = make([]bool, n+1)

	for p := 2; p*p <= n; p++ {
		if !primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = true
			}
		}
	}

	var res []int = make([]int, 0)

	for i := 2; i <= n; i++ {
		if !primes[i] {
			res = append(res, i)
		}
	}

	return res
}
