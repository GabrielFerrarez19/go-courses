package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	slice := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite um numero: ")
	scanner.Scan()

	input := scanner.Text()
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Erro: você deve digitar um número inteiro válido.")
		return
	}

	contain := slices.Contains(slice, num)
	fmt.Println(slices.Contains(slice, 10))
	if contain {
		fmt.Printf("O slice contem %d", num)
	} else {
		fmt.Printf("O slice Não contem %d", num)
	}
}

func Contains[T comparable](s []T, cmp T) bool {
	for _, str := range s {
		if str == cmp {
			return true
		}
	}
	return false
}
