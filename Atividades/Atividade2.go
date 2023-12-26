package main

import "fmt"

func solution(r string) int {
    var result int
    rti := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    for a := 0; a < len(r); a++{
        if len(r)-a > 1{
            if rti[r[a:a+1]] < rti[r[a+1:a+2]] && rti[r[a:a+1]]*10 >= rti[r[a+1:a+2]]{
                result = result-(rti[r[a:a+1]])
            }else{
            result = rti[r[a:a+1]]+result
            }
        }else{
        result = rti[r[a:a+1]]+result
        }
    }
    return result
}
    
func main(){
    // Aqui vocé coloca os inputs da sua maneira preferida
    var test string
    fmt.Printf("Digite um número em algarismos romanos:")
    fmt.Scanf("%s", &test)
	fmt.Printf("O número em algarismos romanos %s no sistema de números índo-arabícos é %d\n", test, solution(test))
}
