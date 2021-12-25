package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	if !validDigit(digit) {
		return ErrDigit
	}

	g[row][column] = digit
	return nil
}

func main() {
	var g Grid
	err := g.Set(0, 0, 15)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de parametres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
