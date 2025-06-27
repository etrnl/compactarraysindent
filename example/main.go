package main

import (
	"fmt"
	"log"

	compactarrays "github.com/etrnl/compactarraysindent"
)

type Game struct {
	Title     string   `json:"title"`
	Platforms []string `json:"platforms"`
	Ratings   []int    `json:"ratings"` // [0, 1, 0, 1, 0] → ESRB flags, etc.
	Genre     string   `json:"genre"`
}

func main() {
	games := []Game{
		{
			Title:     "Chrono Dive",
			Platforms: []string{"PC", "Switch"},
			Ratings:   []int{0, 1, 0, 1, 0},
			Genre:     "Action RPG",
		},
		{
			Title:     "Pixel Racer 2049",
			Platforms: []string{"PC", "PS5"},
			Ratings:   []int{1, 0, 0, 1, 1},
			Genre:     "Arcade",
		},
	}

	// Compact only the noisy arrays
	out, err := compactarrays.CompactMarshalIndent(
		games,
		[]string{"platforms", "ratings"},
		"", "  ",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
