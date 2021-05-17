package GCodeProcessor

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("Function '%s' took %s", name, elapsed)
}

func isValidStartingChar(c rune) bool {
	if (c >= 65 && c <= 96) || (c >= 97 && c <= 122) || c == '(' {
		return true
	}
	return false
}

func isValidValue(c uint8) bool {
	if (c >= 48 && c <= 57) || c == 46 || c == 45 || c == 43 {
		return true
	}
	return false
}

func isCommentStart(c rune) bool {
	if c == '(' {
		return true
	}
	return false
}

func ParseBlock(s string) (GCodeBlock, error) {
	line := GCodeBlock{}
	s = strings.ReplaceAll(s, " ", "")

	// Find, save, and strip out EOL Comment if it exists (starting with ;)
	if i := strings.Index(s, ";"); i >= 0 {
		line.Comment = s[i+1:]
		s = strings.TrimSpace(s[:i])
	}
	// Blank line
	if s == "" {
		return line, nil
	}

	buffer := ""
	for s != "" {
		c := rune(s[0])
		if isValidStartingChar(c) {
			word := GCodeWord{}
			buffer = ""
			if isCommentStart(c) {
				closeIdx := strings.Index(s, ")")
				if closeIdx > 0 {
					word.Comment = s[0 : closeIdx+1]
					line.Words = append(line.Words, word)
					s = s[closeIdx+1:]
				} else {
					return line, errors.New(fmt.Sprintf("Unable to find Comment close ) - ( found at index %d, last ) character found at index %d", 0, closeIdx))
				}
			} else {
				i := 1
				for i = 1; i < len(s); i++ {
					if isValidValue(s[i]) {
						buffer = buffer + string(s[i])
					} else {
						break
					}
				}
				if len(buffer) == 0 {
					return line, errors.New(fmt.Sprintf("Expected word command '%c' to be followed by valid Value, but instead followed by '%c'", s[0], s[1]))
				}
				word.Letter = strings.ToUpper(string(s[0]))
				value, err := strconv.ParseFloat(buffer, 64)
				if err != nil {
					return line, errors.New(fmt.Sprintf("Unable to parse float from Value %s", buffer))
				}
				word.Value = value
				line.Words = append(line.Words, word)
				s = s[i:]
			}
		} else {
			return line, errors.New(fmt.Sprintf("Expected valid word start (alphabetic or Comment start), received %c", c))
		}
	}
	return line, nil
}

func ParseFile(fp *os.File) (*GCodeFile, error) {
	file := &GCodeFile{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line, err := ParseBlock(scanner.Text())
		if err != nil {
			return file, err
		}
		file.Blocks = append(file.Blocks, line)
	}
	return file, nil
}

func ParseGCodeFile(filePath string) (*GCodeFile, error) {
	defer timeTrack(time.Now(), "ParseGCodeFile")
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	file, err := ParseFile(f)
	if err != nil {
		return file, err
	}

	return file, nil
}
