package main

import "fmt"

func solution(nums []int, target int) [2]int {
	for m := 0; m <= len(nums); m++ {
		for n := m + 1; n <= len(nums); n++ {
			if nums[m]+nums[n] == target {
				return [2]int{m, n}
			}
		}
	}
	return [2]int{-1, -2}
}

func main() {
	// Aqui é para estar os inputs
	test := solution([]int{3, 4, 3, -3, -7, -1}, 0)
	fmt.Printf("Lista = %d, Alvo = %d\n", []int{3, 4, 3, -3, -7, -1}, 0)
	fmt.Println("Resultado =", test)
}
