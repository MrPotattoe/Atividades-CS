package main

import "fmt"

var primes []bool = []bool{false, false, true, true, false, true}

func solution(p int) int {
	var res, cont int
	for n := len(primes); p*15 > len(primes); n++ {
		ocont := 0
		for l := 2; l <= n/2+1; l++ {
			if !primes[l] {
				continue
			}
			if n%l == 0 {
				ocont++
				break
			}
		}
		if ocont != 0 {
			primes = append(primes, false)
		}
		if ocont == 0 {
			primes = append(primes, true)
		}
	}
	for y := 3; y <= len(primes); y += 2 {
		if primes[y] {
			cont++
		}
		if cont == p-1 {
			res = y
			break
		}
	}
	return res
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
	a := test{3, 5, solution(3)}
	b := test{7, 17, solution(7)}
	c := test{10, 29, solution(10)}
	d := test{111, 607, solution(111)}
	e := test{345, 2333, solution(345)}
	f := test{12300, 131701, solution(12300)}
	g := test{666, 4973, solution(666)}

	fmt.Printf("Test results\n")
	checker(a)
	checker(b)
	checker(c)
	checker(d)
	checker(e)
	checker(f)
	checker(g)
	overallChecher(testResults)
}
