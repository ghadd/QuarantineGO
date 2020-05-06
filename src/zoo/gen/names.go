package gen

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

const filename = "zoo/res/names.txt"

// GetName gets a name
func GetName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	return lines[r.Intn(len(lines))]
}
