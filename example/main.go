package main

import (
	"fmt"
	"log"

	"github.com/user0608/numeroaletras"
)

func main() {
	numero := numeroaletras.NewNumeroALetras()

	number := 1234567.89
	words, err := numero.ToWords(number, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ToWords:", words)

	money, err := numero.ToMoney(number, 2, "PESOS", "CENTAVOS")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ToMoney:", money)

	str, err := numero.ToString(number, 2, "PESOS", "CENTAVOS")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ToString:", str)

	invoice, err := numero.ToInvoice(number, 2, "PESOS")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ToInvoice:", invoice)
}
