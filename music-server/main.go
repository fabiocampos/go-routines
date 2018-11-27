package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", GetGuitarHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetGuitarHandler(w http.ResponseWriter, r *http.Request) {
	var riff string
	beatChannel := make(chan int, 3)
	go GetBeat(beatChannel)

	for beat := range beatChannel {
		log.Println(beat)
		chord := GetGuitarChordByBeat(beat)
		log.Println(chord)
		riff = fmt.Sprintf("%s %s", riff, chord)
	}

	content, _ := json.Marshal(map[string]string{"Guitar": riff})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func GetGuitarChordByBeat(beat int) string {
	chordsKnown := []string{"D", "F#", "E"}
	return chordsKnown[beat]
}

func GetBeat(beatChannel chan<- int) {
	for i := 0; i < cap(beatChannel); i++ {
		time.Sleep(1 * time.Second)
		chordIndex := rand.Intn(3)
		beatChannel <- chordIndex
	}
	close(beatChannel)
}
