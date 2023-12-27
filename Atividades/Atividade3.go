package main

import "fmt"

func solution(x int) []int {
	var d []int
	if x > 0 {
		for i := 1; i <= x; i++ {
			if x%i == 0 {
				d = append(d, i)
			}
		}
	}
	if x < 0 {
		xx := -1 * x
		for i := 1; i <= xx; i++ {
			if xx%i == 0 {
				d = append(d, i)
			}
		}
	}
	return d
}

func main() {
	var x int
	fmt.Printf("Digite um número: ")
	fmt.Scanf("%d", &x)
	divs := solution(x)
	fmt.Printf("O(s) divisor(es) é(são) ")
	for i := 0; i < len(divs); i++ {
		fmt.Printf("%d ", divs[i])
	}
	fmt.Println("")
}
