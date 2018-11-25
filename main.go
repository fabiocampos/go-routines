package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		PlayGuitar()
		wg.Done()
	}()
	go func() {
		PlayDrums()
		wg.Done()
	}()
	wg.Wait()
}

func PlayGuitar() {
	riffs := []string{"E G A", "E D C", "D D E"}
	for _, riffToPerform := range riffs {
		fmt.Println(riffToPerform)
		time.Sleep(500 * time.Millisecond)
	}
}

func PlayDrums() {
	beats := []string{"CC1 BD CC2 BD", "CC2 BD CC1 DB", "CC1 LMT LT HFT"}
	for _, beatToPerform := range beats {
		fmt.Println(beatToPerform)
		time.Sleep(250 * time.Millisecond)
	}
}
