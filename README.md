# :black_joker: baccarat [![CircleCI](https://circleci.com/gh/bigUNO/baccarat/tree/master.svg?style=svg)](https://circleci.com/gh/bigUNO/baccarat/tree/master)

A simple package that provides baccarat functionality.

## Installation

```
go get -u github.com/bigUNO/baccarat
```

## Quick Start

```golang
// Setup a new game of baccarat.
game, err := baccarat.New()
if err != nil {
  panic(err) // handle error
}
fmt.Printf("Starting game:%v\n", game.ID)

// Deal the first hand
game.Start()

// Print the score and clear current cards
game.CloseHand()
```

## Development Status: Unstable

Breaking changes are coming soon.
