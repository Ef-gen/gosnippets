// Package rfkb provides read different type of values from the keyboard
package rfkb

import (
	"bufio"
	"log"
	"strings"
)
import "os"

// ReadString reads string types of values
// input ends with a line break
// deleting '\n' rune and return inputed string. 
func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(str)
}
