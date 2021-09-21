package rfkb

import (
	"bufio"
	"log"
	"strings"
)
import "os"

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(str)
}
