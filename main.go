package main

import (
	"fmt"
	"time"
)

func main() {
	guitarRiffChannel := make(chan string)
	drumBeatChannel := make(chan string)
	go PlayGuitar(guitarRiffChannel)
	go PlayDrums(drumBeatChannel)
	for {
		select {
		case guitarrRiff, isGuitarChannelOpen := <-guitarRiffChannel:
			if isGuitarChannelOpen == true {
				fmt.Println(guitarrRiff)
			}
		case drumBeat, isDrumChannelOpen := <-drumBeatChannel:
			if isDrumChannelOpen == true {
				fmt.Println(drumBeat)
			}
		case <-time.After(2 * time.Second):
			close(guitarRiffChannel)
			close(drumBeatChannel)
			fmt.Println("######Aplauses!!!!!")
			return
		}
	}
}

func PlayGuitar(riff chan<- string) {
	riffs := []string{"E G A", "E D C", "D D E"}
	for _, riffToPerform := range riffs {
		riff <- riffToPerform
		time.Sleep(500 * time.Millisecond)
	}
}

func PlayDrums(beat chan<- string) {
	beats := []string{"CC1 BD CC2 BD", "CC2 BD CC1 DB", "CC1 LMT LT HFT"}
	for _, beatToPerform := range beats {
		beat <- beatToPerform
		time.Sleep(250 * time.Millisecond)
	}
}
