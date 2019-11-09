package main

import (
	"fmt"

	"github.com/bigUNO/baccarat"
)

func main() {
	game, err := baccarat.New()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Starting game:%v\n", game.ID)
	for i := 0; i < 10; i++ {
		fmt.Printf("Round %v> ", i)
		game.Start()
		game.CloseHand()
	}
}
