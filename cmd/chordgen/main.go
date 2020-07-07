package main

import (
	"flag"
	"github.com/main-kun/chordgen/pkg/picgen"
	"log"
	"strconv"
	"strings"
)

const maxKeysIndex int = 23

func main() {
	highlightUsage := "Comma separated list of highlighted keys\nkeys are indexed by semitones from 0 to 23 eg C=0, C#=1, D=2 ..."

	highlightsPtr := flag.String("h", "", highlightUsage)
	filenamePtr := flag.String("f", "output.png", "output file name")

	flag.Parse()

	if *highlightsPtr == "" {
		log.Fatal("-h highlights list is required")
	}
	highlightsArr := strings.Split(*highlightsPtr, ",")
	var highlightsIntArr []int
	const intSize = 32 << (^int(0) >> 32 & 1)
	for _, item := range highlightsArr {
		parsedItem, err := strconv.ParseInt(item, 10, intSize)
		parsedInt := int(parsedItem)
		if err != nil {
			log.Fatal("failed to parse highlights array")
		}
		if parsedInt > maxKeysIndex || parsedInt < 0 {
			log.Fatal("highlights index out of bounds")
		}
		highlightsIntArr = append(highlightsIntArr, parsedInt)
	}

	err := picgen.DrawChord(highlightsIntArr, *filenamePtr)
	if err != nil {
		log.Fatal(err)
	}
}
