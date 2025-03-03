package deferpanic

import (
	"fmt"
	"log"
)

// instruccion de go que permite configurar cual va a ser la ultima instrucción
// en ejecutarse cuando salga de la función
func VemosDefer() {
	fmt.Println("Este es el primer mensaje")

	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")
}

// Panic isntrucción que aborta un programa con un mensaje que lo envio a consola
func EjemploPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1

	if a == 1 {
		panic("Se encontro el valor 1")
	}

}
