package main

import (
	"Assignment1/TicTacToeStruct/game"
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered from", r)
		}
	}()

	var result string
	var err error

	game1 := game.NewGame("Ankush", "Sanjeev")
	// fmt.Println(game1.Players[0])
	// fmt.Println(game1.Players[1])

	result, err = game1.Play(0)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	result, err = game1.Play(5)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result, err = game1.Play(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result, err = game1.Play(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result, err = game1.Play(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// game1.Play(0)
	// game1.Play(5)
	// game1.Play(1)
	// game1.Play(4)
	// game1.Play(2)
	// game1.Play(3)

}
