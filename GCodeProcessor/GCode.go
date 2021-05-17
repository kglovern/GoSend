package GCodeProcessor

import (
	"fmt"
	"log"
	"strings"
)

// GCodeWord
// A Single G-Code consists of a Letter command with a corresponding Value and an optional trailing Comment
type GCodeWord struct {
	Letter  string
	Value   float64
	Comment string
}

// GCodeBlock
// A Single GCodeBlock consists of 1 or more G-Code Words
type GCodeBlock struct {
	Words   []GCodeWord
	Comment string
}

type GCodeFile struct {
	Blocks        []GCodeBlock
	Tools         []GCodeWord
	UnitModals    []GCodeWord
	SpindleSpeeds []GCodeWord
	FeedRates     []GCodeWord
}

// String handlers for structs
func (l *GCodeBlock) String() string {
	var s strings.Builder
	for _, word := range l.Words {
		_, err := fmt.Fprintf(&s, "%s ", word.String())
		if err != nil {
			log.Fatal(err)
		}
	}
	return strings.TrimSpace(s.String())
}

func (c *GCodeWord) String() string {
	if c.Comment != "" {
		return ""
	}
	return c.Letter + strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", c.Value), "0"), ".")
}

func (f *GCodeFile) String() string {
	var s strings.Builder
	for _, line := range f.Blocks {
		_, err := fmt.Fprintf(&s, "%s\n", line.String())
		if err != nil {
			log.Fatal(err)
		}
	}
	return s.String()
}
