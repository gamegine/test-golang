package main

import (
	"fmt"
	"sync"
	"time"
)

func task(str string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("task", str)
}

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

func main() {
	wg.Add(1)
	go task("a")
	wg.Add(1)
	go task("b")
	wg.Add(1)
	go task("c")

	fmt.Println("tasks started")
	wg.Wait()
	fmt.Println("end")
}
