package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

func main() {
	s := starship{laser(3)}

	fmt.Println(s.talk())
	shout(s)
}
