package eje_rinterfaces

import (
	"fmt"

	"github.com/rsanchezs151/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())

}
