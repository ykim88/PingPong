package gamer

import (
	"fmt"
	"sync"
)

type Gamer interface {
	Play(ballChannel chan int)
	StopGame()
	Named(name string)
}

type gamer struct {
	stop       chan bool
	endWork    *sync.WaitGroup
	name       string
	finishGame bool
}

func GetGamer(wg *sync.WaitGroup) Gamer {
	gamer := new(gamer)
	gamer.finishGame = false
	gamer.setEndWork(wg)
	return gamer
}

func (g *gamer) Named(name string) {
	g.name = name
}

func (g *gamer) StopGame() {
	g.finishGame = true
}

func (g *gamer) isFinishGame() bool {
	return g.finishGame
}

func (g *gamer) setEndWork(wg *sync.WaitGroup) {
	g.endWork = wg
	wg.Add(1)
}

func (g *gamer) Play(ballChannel chan int) {
	defer g.endWork.Done()

	for !g.isFinishGame() {
		ball := <-ballChannel
		fmt.Printf("%s: %d\n", g.name, ball)
		ball++
		ballChannel <- ball
	}
}
