package main

import (
	"fmt"

	"github.com/rsanchezs151/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(65465)

	fmt.Println(estado)
	fmt.Println(texto)

}
