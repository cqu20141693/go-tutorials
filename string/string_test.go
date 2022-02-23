package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	logDir1 := ""
	logDir2 := "log"
	logDir3 := "/var/log"

	checkDir(logDir1)
	checkDir(logDir2)
	checkDir(logDir3)
}

func checkDir(dir string) {
	if dir == "" {
		fmt.Println("empty")
	} else if strings.Contains(dir, "/") {
		fmt.Println("contains /")
	} else {
		fmt.Println("else")
	}

}
