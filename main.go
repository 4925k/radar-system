package main

import "radar-system/speed"

func main() {
	speeds := []float64{60.4254, 60.43, 60.48, 61.58, 61.32, 61.58}
	var roundUpUnit float64 = 0.05
	speedCount := make(map[float64]int)

	for _, s := range speeds {
		temp := speed.RoundUP(s, roundUpUnit)
		speedCount[temp]++
	}

	speed.Render(&speedCount, roundUpUnit)
}
