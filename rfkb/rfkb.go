// Package rfkb provides read different type of values from the keyboard
package rfkb

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)
import "os"

// ReadString reads string types of values
// input ends with a line break
// deleting '\n' rune and return inputted string.
func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(str)
}

// ReadInt reads string types of values, that may be converted to integer
// input ends with a line break
// deleting '\n' rune, convert to integer and return integer.
func ReadInt() int {
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	newInt, err := strconv.Atoi(strings.TrimSpace(i))
	if err != nil {
		log.Fatal(err)
	}
	return newInt
}

// ReadFloat reads string types of values, that may be converted to float
// input ends with a line break
// deleting '\n' rune, convert to integer and return float.
func ReadFloat() float64 {
	reader := bufio.NewReader(os.Stdin)
	f, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	newFloat, err := strconv.ParseFloat(strings.TrimSpace(f), 64)
	if err != nil {
		log.Fatal(err)
	}
	return newFloat
}
