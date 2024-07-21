# :black_joker: baccarat [![CircleCI](https://dl.circleci.com/status-badge/img/gh/bigUNO/baccarat/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/bigUNO/baccarat/tree/main)

A simple package that provides baccarat functionality.

## Installation

```
go install github.com/bigUNO/baccarat@latest
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

## Development Status: Stable-ish
