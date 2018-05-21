package main

import (
	"GameGoRoutine/entity"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	game := entity.SetNumberOfGamers(5)
	start := time.Now()
	game.StartGame()
	Stop(game)
	fmt.Println(time.Since(start))
}

func Stop(game entity.Game) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	game.StopGame()
}
