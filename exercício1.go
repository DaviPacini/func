//Crie uma função que receba dois parâmetros (a e b) e retorne a divisão entre eles. Caso o
//segundo parâmetro seja zero, retorne um erro.

package main

import "fmt"

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("não é possível dividir por 0")
	} else {
		return a / b, nil
	}
}
func main() {
	var x int
	var y int

	fmt.Println("Informe 2 números para efetuar a divisão: ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	resultado, err := dividir(x, y)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resultado)
	}

}
