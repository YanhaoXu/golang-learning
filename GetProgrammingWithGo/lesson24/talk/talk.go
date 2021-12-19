package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martain struct {
}

func (m martain) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", 3)
}

func main() {
	t = martain{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}
