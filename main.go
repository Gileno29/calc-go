package main

import (
	"fmt"
	"strconv"
)

func sum(x, y float64) float64 {

	return x + y
}

func subtracao(x, y float64) float64 {

	return x - y
}

func multiplicacao(x, y float64) float64 {

	return x * y
}

func divisao(x, y float64) float64 {

	return x / y
}

func main() {

	var choice string
	fmt.Println("Escolha uma opção para realizar a operação")
	fmt.Println("1 - SOMA")
	fmt.Println("2 - SUBTRAÇÃO")
	fmt.Println("3 - MULTIPLICACAO")
	fmt.Println("4 - DIVISAO")
	fmt.Println("0 - Sair")

	for {
		fmt.Scan(&choice)
		choiceInt, err := strconv.ParseFloat(choice, 64)

		if err != nil {
			fmt.Println("Erro ao converter entrada")
			return
		}

		switch choiceInt {

		case float64(1):
			var v1 string
			var v2 string

			fmt.Scan(&v1)
			fmt.Scan(&v2)

			v1Float, err := strconv.ParseFloat(v1, 64)

			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			v2Float, err := strconv.ParseFloat(v2, 64)
			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			fmt.Println("Soma dos valores: ", sum(v1Float, v2Float))

		case float64(2):
			var v1 string
			var v2 string

			fmt.Scan(&v1)
			fmt.Scan(&v2)

			v1Float, err := strconv.ParseFloat(v1, 64)

			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			v2Float, err := strconv.ParseFloat(v2, 64)
			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			fmt.Println("subtracao dos valores: ", subtracao(v1Float, v2Float))

		case 3:
			var v1 string
			var v2 string

			fmt.Scan(&v1)
			fmt.Scan(&v2)

			v1Float, err := strconv.ParseFloat(v1, 64)

			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			v2Float, err := strconv.ParseFloat(v2, 64)
			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			fmt.Println("Multiplicacao dos valores: ", multiplicacao(v1Float, v2Float))

		case 4:
			var v1 string
			var v2 string

			fmt.Scan(&v1)
			fmt.Scan(&v2)

			v1Float, err := strconv.ParseFloat(v1, 64)

			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			v2Float, err := strconv.ParseFloat(v2, 64)
			if err != nil {
				fmt.Println("Erro ao converter entrada da soma")

			}

			fmt.Println("Divisao dos valores: ", divisao(v1Float, v2Float))

		case float64(0):
			fmt.Println("Saindo")
			return

		default:
			fmt.Println("OPÇÃO INEXISTENTE")
		}
	}
}
