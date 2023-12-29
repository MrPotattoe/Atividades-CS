package main

import "fmt"

// Essa foi minha primeira tentativa para a questão 4
// Haverá breves comentários sobre como ela funciona

func solution(n int) int {
// Primeiro eu criei um slice com os dois primeiros primos
	primes := []int{2, 3}
//Grosso modo, esse primeiro looping pega o último primo da lista e, como todos os primos maiores que 2 são ímpares, ele soma dois para verificar se o próximo ímpar é primo. Esse looping se repete até o último índice dessa lista de primos ser o que a gente deseja. Eu enxergo essa ideia de só usar os ímpares como uma otimização
	for i := primes[len(primes)-1] + 2; n > len(primes); i += 2 {
// Esse Contador reflete um pouco minha incompetência de verificar quantos divisores certo múmero tem
		cont := 0
// Agora, esse looping vai verificar se o ímpar "i" é primo. Para isso, ele divide esse "i" por cada primo já presente na lista. Note que isso depende do fato da lista estar sendo formada em ordem crescente
		for j := 1; j < len(primes); j++ {
// Essa condição é uma otimização, pois o loop só verifica os primos que são menores que a metade de "i", já que nenhum número maior que i/2 pode dividir i.
			if primes[j] > i+1/2 {
				break
			}
// Se um dos primos da lista dividir "i", se adiciona +1 ao Contador. Então, como um número primo só tem 1 e ele mesmo de divisores, a condição abaixo para esse looping. Eu também considero para nesse primeiro divisor uma otimização, ao invés de ver toda a lista.
			if i%primes[j] == 0 {
				cont++
			}
			if cont > 0 {
				break
			}
		}
// Por fim, se o último looping tiver sido parado pela terceira condição, então o contador é diferente de 0 e o looping segue para a próxima interação. Caso ele tenha sido interrompido pela primeira condição, então o contador é zero e ele é um número primo
		if cont == 0 {
			primes = append(primes, i)
		}

	}
// O looping mais externo para quando o tamanho da lista de primos for igual ao índice exigido pelo usuário. Assim, como a lista foi sendo criada em ordem crescente, ele imprime o último índice.
	return primes[n-1]
}

func main() {
	var a int
	fmt.Printf("Digite o índice do n-ésimo primo que você quer descobrir: ")
	fmt.Scanf("%d", &a)
	b := solution(a)
	fmt.Printf("O %d-ésimo primo é: %d\n", a, b)
}
