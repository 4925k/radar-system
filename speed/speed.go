package speed

import (
	"fmt"
	"math"
	"strconv"
)

// RoundUp rounds up the value up the to roundUpUnit
// e.g 6.4267 rounded by 0.05 is 6.45
func RoundUP(value, roundUpUnit float64) float64 {
	return value + roundUpUnit - math.Mod(value, roundUpUnit)
}

func Render(speedCount *map[float64]int, roundUpUnit float64) {
	fmt.Println("Vehicle Speed\tCount")
	for k, v := range *speedCount {
		fmt.Printf("%skm/h\t%d\n", strconv.FormatFloat(k, 'f', 2, 64), v)
	}
}
