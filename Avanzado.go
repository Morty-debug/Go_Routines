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
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Start"

	time.Sleep(1 * time.Second)
	command <- "Shutdown"

	wg.Wait()
}

