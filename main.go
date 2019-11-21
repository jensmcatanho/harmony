package main

import (
	"fmt"
	"regexp"
)

var (
	greek = map[int]string{
		0: "IM",
		1: "IIm",
		2: "IIIm",
		3: "IVM",
		4: "VM",
		5: "VIm",
		6: "VIIo",
	}
)

type Chord struct {
	Name string
}

func NewChord(name string) *Chord {
	return &Chord{
		Name: name,
	}
}

func main() {
	greekPatterns := []string{`([A-G][M])`, `([A-G][m])`, `([A-G][m])`, `([A-G][M])`, `([A-G][M])`, `([A-G][m])`, `([A-G][o])`}

	chord := NewChord("Em7")
	var possibilities []string

	for i, pattern := range greekPatterns {
		match, _ := regexp.MatchString(pattern, chord.Name)

		if match {
			possibilities = append(possibilities, greek[i])
		}
	}

	fmt.Printf("%s's possibilities: %v", chord, possibilities)
}
