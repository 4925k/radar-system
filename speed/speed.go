package speed

import (
	"fmt"
	"math"
	"strings"
)

type Speeds map[float64]int
type record struct {
	Speed string `json:"speed"`
	Count int    `json:"count"`
}

// RoundUp rounds up the value up the to roundUpUnit
// e.g 6.4267 rounded by 0.05 is 6.45
func RoundUP(value, roundUpUnit float64) float64 {
	return value + roundUpUnit - math.Mod(value, roundUpUnit)
}

// RenderCLI prints out a table of the current speeds and their count to the CLI
func (s *Speeds) RenderCLI(roundUpUnit float64) {
	fmt.Println("Vehicle Speed\tCount")
	for k, v := range *s {
		fmt.Printf("%.*fkm/h\t%d\n", decimalCount(roundUpUnit), k, v)
	}
}

// JSON returns with a json structure of the current speeds
func (s *Speeds) RenderJSON(roundUpUnit float64) []record {
	var speeds []record

	for k, v := range *s {
		speeds = append(speeds, record{fmt.Sprintf("%.*f", decimalCount(roundUpUnit), k), v})
	}

	return speeds
}

func decimalCount(value float64) int {
	return len(strings.Split(fmt.Sprintf("%v", value), ".")[1])
}
