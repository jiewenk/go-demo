package main

import (
	"fmt"
	"time"
)

type MVIPlayer struct {
	stat    int
	process int
}

func (p MVIPlayer) Play(source string) {
	fmt.Println("Playing MVI music", source)
	p.process = 0
	for p.process < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.process += 10
	}
	fmt.Println("\n Finished playing", source)
}
