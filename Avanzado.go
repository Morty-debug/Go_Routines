package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Start"
	for {
		select {
			case cmd := <-command:
				fmt.Println(cmd)
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

func main() {
	/* Invocamos Go Routine */
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg) //se ejecuta por defecto en Start
	
	time.Sleep(1 * time.Second) //esperar un segundo
	command <- "Pause" //despues de un segundo pausamos la rutina

	time.Sleep(1 * time.Second) //esperar un segundo
	command <- "Start" //despues de un segundo ejecutamos nuevamente la rutina donde se quedo

	time.Sleep(1 * time.Second) //esperar un segundo
	command <- "Shutdown" //despues de un segundo detenemos el proceso

	wg.Wait() //finalizamos rutina
}

