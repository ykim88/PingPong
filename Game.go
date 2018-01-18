package main

import (
	"fmt"
	"time"
)

func play(ballChannel <-chan int, out chan<- int, entity string) {
	for {
		ball := <-ballChannel
		fmt.Printf("%s turn %d\n", entity, ball)
		ball++
		//ballChannel <- ball
		out <- ball
	}
}

func startGame(gamer1 string, gamer2 string) {
	ballChannel1 := make(chan int)
	ballChannel2 := make(chan int)

	ball := 0
	quit := make(chan bool)
	go func(chan bool) {
		fmt.Scanln()
		quit <- false
	}(quit)

	go play(ballChannel1, ballChannel2, gamer1)
	go play(ballChannel2, ballChannel1, gamer2)

	start := time.Now()
	ballChannel1 <- ball

	<-quit
	fmt.Println(time.Now().Sub(start))
}

func main() {
	gamer1 := "Gamer1"
	gamer2 := "Gamer2"

	startGame(gamer1, gamer2)
}
