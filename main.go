package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//variable de retorno
	retorno := make(chan string, 1)
	//proceso en paralelo
	go func() {
		texto := proceso() //proceso que puede tardar demasiado
		retorno <- texto
	}()
	//matar proceso con un tiempo limite o por retornar datos si no excede el tiempo
	select {
		case resultado := <-retorno: fmt.Println(resultado)
		case <-time.After(3 * time.Second): fmt.Println("[ERROR] tiempo expirado, proceso muerto")
	}
}

func proceso() string {
	rand.Seed(time.Now().UTC().UnixNano())  //ramdon es seteado por el tiempo del sistema
	naleatorio := 1 + rand.Intn(5-1) 	//numero aleatorio entre 1 y 5
	fmt.Println("Tiempo en que se ejecutara: ",naleatorio," segundos")
	for {
		if naleatorio <= 0 {
			break
		} else {
			fmt.Println(naleatorio)		//muestra que sigue esperando
			time.Sleep(1 * time.Second)  
			naleatorio--
		}
	}
	return "[OK] proceso finalizado"
}
