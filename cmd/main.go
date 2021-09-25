package main

import (
	"fmt"
	"quizz/internal"
	"strings"
)

var example_csv = strings.NewReader(`5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7`)

func main() {
	game := internal.NewCLIGame()
	if game.Play(example_csv) {
		fmt.Println("BRAVO! 👑")
		return
	}
	fmt.Println("You lost  ☹️")
}
