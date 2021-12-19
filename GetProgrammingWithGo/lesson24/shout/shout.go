package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martain struct {
}

func (m martain) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martain{})
	shout(laser(5))
}
