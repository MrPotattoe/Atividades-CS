package main

import "fmt"

// Esaa é uma limitação desse código, você tem ter de antemão uma boa estimativa do tamanho do número que você quer
const MAX_SIZE = 500000000

var primes []int
var IsPrime [MAX_SIZE]bool

func SieveOfEratosthenes() {
	for x := 2; x < len(IsPrime); x++ {
		IsPrime[x] = true
	}
	for a := 2; a*a < len(IsPrime); a++ {
		if IsPrime[a] {
			primes = append(primes, a)
			for k := a * a; k < MAX_SIZE; k += a {
				IsPrime[k] = false
			}
		}
	}
}

func solution(n int) int {
	return primes[n-1]
}

// A partir daquu tem somente um método, uma estrutura e uma função para testar os resultados, além da função principal.
type test struct {
	in, expect, out int
}

var testResults []bool

func checker(t test) {
	if t.expect == t.out {
		fmt.Printf("Input: %d\tExpected: %d\tResult: Paassed\n", t.in, t.expect)
		fmt.Printf("==================================================================\n")
		testResults = append(testResults, true)
	}
	if t.expect != t.out {
		fmt.Printf("Input: %d\tExpected: %d\tResult: Reproved\n", t.in, t.expect)
		fmt.Printf("Errors: Expected (%d) != Output (%d)\n", t.expect, t.out)
		fmt.Printf("==================================================================\n")
		testResults = append(testResults, false)
	}
}

func overallChecher([]bool) {
	anocont := 0
	for _, q := range testResults {
		if !q {
			anocont++
			break
		}
	}
	if anocont != 0 {
		fmt.Printf("\n")
		fmt.Printf("===================================================\n")
		fmt.Println("\tFINAL RESULT OF EXERCISE: FAILED")
		fmt.Printf("===================================================\n")
	}
	if anocont == 0 {
		fmt.Printf("\n")
		fmt.Printf("===================================================\n")
		fmt.Println("\tFINAL RESULT OF EXERCISE: SUCCESS")
		fmt.Printf("===================================================\n")
	}
}

func main() {
	SieveOfEratosthenes()

	a := test{1, 2, solution(1)}
	b := test{2, 3, solution(2)}
	c := test{3, 5, solution(3)}
	d := test{5, 11, solution(5)}
	e := test{13, 41, solution(13)}

	fmt.Println("Test Results")
	checker(a)
	checker(b)
	checker(c)
	checker(d)
	checker(e)
	overallChecher(testResults)
}
