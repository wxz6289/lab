package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string) {
	fmt.Println("Shopping...", name)
	time.Sleep(2 * time.Second)
	fmt.Println("Shopping done.", name)
	waitGroup.Done()
}

var waitGroup sync.WaitGroup

func show() {
	start := time.Now()
	shopping("Alice")
	shopping("Bob")
	shopping("Charlie")
	fmt.Println("Total time:", time.Since(start))
}

func show_goroutine() {
	start := time.Now()
	waitGroup.Add(3)
	go shopping("Alice")
	go shopping("Bob")
	go shopping("Charlie")
	// time.Sleep(3 * time.Second)
	waitGroup.Wait()
	fmt.Println("Total time:", time.Since(start))
}

var payChannel chan int = make(chan int, 1)

func pay(name string, price int, wait *sync.WaitGroup) {
	fmt.Println("Paying...", name, price)
	time.Sleep(2 * time.Second)
	fmt.Println("Payment done.", name, price)
	payChannel <- price
	wait.Done()
}

func show_pay() {
	start := time.Now()
	wait := sync.WaitGroup{}
	wait.Add(3)
	go pay("Alice", 100, &wait)
	go pay("Bob", 200, &wait)
	go pay("Charlie", 300, &wait)
	total := 0
	go func() {
		defer close(payChannel)
		wait.Wait()
	}()

	for price := range payChannel {
		total += price
	}
	fmt.Println("Total time:", time.Since(start))
	fmt.Println("Total payment:", total)
}

var nameChannel chan string = make(chan string, 1)
var priceChannel chan int = make(chan int, 1)

func send(name string, price int, wait *sync.WaitGroup) {
	fmt.Println("Sending...", name, price)
	time.Sleep(2 * time.Second)
	fmt.Println("Sending done.", name, price)
	nameChannel <- name
	priceChannel <- price
	wait.Done()
}

func show_send() {
	wait := sync.WaitGroup{}
	wait.Add(3)
	start := time.Now()
	go send("Alice", 100, &wait)
	go send("Bob", 200, &wait)
	go send("Charlie", 300, &wait)
	go func() {
		wait.Wait()
		close(nameChannel)
		close(priceChannel)
	}()

	fmt.Println("Total time:", time.Since(start))

	nameList := []string{}
	priceList := []int{}

	event := func() {
		nameClosed := false
		priceClosed := false
		for !nameClosed || !priceClosed {
			select {
			case name, ok := <-nameChannel:
				if ok {
					nameList = append(nameList, name)
				} else {
					nameClosed = true
				}
			case price, ok := <-priceChannel:
				if ok {
					priceList = append(priceList, price)
				} else {
					priceClosed = true
				}
			}
		}
		fmt.Println("All done.")
	}
	event()
	fmt.Println("Names:", nameList)
	fmt.Println("Prices:", priceList)
}

var done = make(chan bool, 1)

func see() {
	fmt.Println("Starting ...")
	time.Sleep(4 * time.Second)
	fmt.Println("Done.")
	done <- true
}

func show_see() {
	go see()
	select {
	case <-done:
		fmt.Println("see done.")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}

var total int = 0
var totalWait sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock()
	for i := 0; i < 1000; i++ {
		total++
	}
	lock.Unlock()
	totalWait.Done()
}

func sub() {
	lock.Lock()
	for i := 0; i < 1000; i++ {
		total--
	}
	lock.Unlock()
	totalWait.Done()
}

func show_total() {
	totalWait.Add(2)
	go add()
	go sub()
	totalWait.Wait()
	fmt.Println("Total:", total)
}


var nameMap = make(map[string]int)


func show_map() {
	go func() {
		for {
			nameMap["Alice"] = 100
		}
	}()

	go func() {
		for {
		fmt.Println("Alice's price:", nameMap["Alice"])
		}
	}()

	select {}
}

func safe_map_show() {
	var safeMap sync.Map

	go func() {
		for {
			safeMap.Store("Alice", 100)
		}
	}()

	go func() {
		for {
			value, ok := safeMap.Load("Alice")
			if ok {
				fmt.Println("Alice's price:", value)
			}
		}
	}()

	select {}
}
