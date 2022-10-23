package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func main() {
	fmt.Println("Ejemplo1, basado en procesos que no tienen definido su tiempo de ejecucion y para evitar esperas infinitas definimos un temporizador de 3 segundos:")
	ejemplo1()
	fmt.Println("\nEjemplo2, basado en procesos que se ejecutan en paralelo del resto de la aplicacion. Pero definimos controles para poder pausarlo y detenerlo:")
	ejemplo2()
}



/*****************************************************************/
/* 			Inicio ejemplo1 			 */
/*****************************************************************/
func ejemplo1() {
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



/*****************************************************************/
/* 			Inicio ejemplo2 			 */
/*****************************************************************/
var i int
func ejemplo2() {
	/* Invocamos Go Routine */
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)

	fmt.Println("Start routine")
	go routine(command, &wg) //se ejecuta por defecto en Start
	time.Sleep(1 * time.Second) //esperar un segundo y mientras esperamos el proceso en paralelo sigue ejecutandose

	command <- "Pause" //pausamos la rutina
	fmt.Println("Pause routine")
	time.Sleep(1 * time.Second) //esperar un segundo y mientras esperamos el proceso en paralelo sigue pausado

	fmt.Println("Start routine") 
	command <- "Start" //ejecutamos nuevamente la rutina donde se quedo
	time.Sleep(1 * time.Second) //esperar un segundo y mientras esperamos el proceso en paralelo sigue ejecutandose

	command <- "Shutdown" //detenemos el proceso
	fmt.Println("Shutdown routine")
	
	wg.Wait() //finalizamos rutina
}
func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Start"
	for {
		select {
			case cmd := <-command:
				//fmt.Println(cmd)
				switch cmd {
					case "Shutdown": return
					case "Pause": status = "Pause"
					default: status = "Start"
				}
			default:
				if status == "Start" {
					work()
				}
		}
	}
}
func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}
