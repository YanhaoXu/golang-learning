package main

import (
	"image"
	"log"
	"time"
)

const (
	right = command(0)
	left  = command(1)
)

type command int

type RoverDiver struct {
	commandc chan command
}

func (r *RoverDiver) Left() {
	r.commandc <- left
}

func (r *RoverDiver) Right() {
	r.commandc <- right
}

func (r RoverDiver) driver() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func NewRoverDiver() *RoverDiver {
	r := &RoverDiver{
		commandc: make(chan command),
	}
	go r.driver()
	return r
}

func main() {
	r := NewRoverDiver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
