package main

import (
	//"github.com/rsanchezs151/godesde0/users"
	//e "github.com/rsanchezs151/godesde0/defer_panic"
	//"fmt"

	//e "github.com/rsanchezs151/godesde0/goroutines"
	//e "github.com/rsanchezs151/godesde0/webserver"
	e "github.com/rsanchezs151/godesde0/middelware"
)

func main() {
	// estado, texto := variables.ConviertoaTexto(65465)

	// fmt.Println(estado)
	// fmt.Println(texto)
	/*
				Pedro := new(models.Hombre)
				e.HumanosRespirando(Pedro)
				Ana := new(models.Mujer)
				e.HumanosRespirando(Ana)

			e.VemosDefer()
			e.EjemploPanic()
			canal1 := make(chan bool)
		go e.MiNombreLento("Roberto", canal1)
		fmt.Println("Estoy aqui")
		defer func() {
			<-canal1
		}()

	*/
	e.MiMiddelware()
}
