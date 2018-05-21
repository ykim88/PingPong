package entity

import (
	"GameGoRoutine/entity/gamer"
	"fmt"
	"sync"
)

type Game interface {
	StartGame()
	StopGame()
}

type game struct {
	Gamers      []gamer.Gamer
	wg          sync.WaitGroup
	ballChannel chan int
	stop        chan bool
}

func SetNumberOfGamers(number int) Game {
	g := new(game)
	for i := 0; i < number; i++ {
		g.Gamers = append(g.Gamers, gamer.GetGamer(&g.wg))
	}
	return g
}

func (g *game) StopGame() {
	defer close(g.ballChannel)
	for _, player := range g.Gamers {
		player.StopGame()
	}
	g.wg.Wait()
}

func (g *game) StartGame() {
	gamerNumber := len(g.Gamers)
	g.ballChannel = make(chan int, (gamerNumber/2)+(gamerNumber%2))

	ball := 0
	g.ballChannel <- ball
	g.gamePlay()
}

func (g *game) gamePlay() {

	for i, gamer := range g.Gamers {
		gamer.Named(fmt.Sprintf("Gamer%d", i))
		go gamer.Play(g.ballChannel)
	}
}
