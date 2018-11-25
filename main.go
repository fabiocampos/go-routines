package main

import (
	"fmt"
	"time"
)

func main() {
	PlayGuitar("E A D G B e")
	PlayDrums("CC1 BD CC2 BD")
}

func PlayGuitar(riff string) {
	fmt.Println(riff)
	time.Sleep(600 * time.Millisecond)
}

func PlayDrums(beat string) {
	fmt.Println(beat)
	time.Sleep(300 * time.Millisecond)
}
