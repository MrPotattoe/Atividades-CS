package main

import "fmt"

// Essa foi minha primeira ideia para criar um "cache"
var primes []int = []int{2, 3}

func solution(n int) int {
	for i := primes[len(primes)-1] + 2; n > len(primes); i += 2 {
		cont := 0
		for j := 1; j < len(primes); j++ {
			if primes[j] > i+1/2 {
				break
			}
			if i%primes[j] == 0 {
				cont++
			}
			if cont > 0 {
				break
			}
		}
		if cont == 0 {
			primes = append(primes, i)
		}

	}
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
	a := test{183739, 2509849, solution(183739)}
	b := test{678, 5077, solution(678)}
	c := test{83839, 1074107, solution(83839)}
	d := test{72, 359, solution(72)}
	e := test{777, 5903, solution(777)}
	f := test{4, 7, solution(4)}

	checker(a)
	checker(b)
	checker(c)
	checker(d)
	checker(e)
	checker(f)
	overallChecher(testResults)
}
