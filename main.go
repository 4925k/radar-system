package main

import (
	"fmt"
	"radar-system/server"
	"radar-system/speed"
	"strconv"
	"sync"
)

//need to do
// format string, count decimal places

var ()

func main() {
	var wg sync.WaitGroup
	listener := make(chan []byte)

	wg.Add(2)
	go func() {
		server.Listen(listener)
		wg.Done()
	}()

	go func() {
		for data := range listener {
			f, err := strconv.ParseFloat(string(data), 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%f\n", f)
		}
		wg.Done()
	}()

	speeds := []float64{60.4254, 60.43, 60.48, 61.58, 61.32, 61.58}
	var roundUpUnit float64 = 0.05

	speedCount := make(speed.Speeds)

	for _, s := range speeds {
		temp := speed.RoundUP(s, roundUpUnit)
		speedCount[temp]++
	}

	speedCount.RenderCLI(roundUpUnit)

	fmt.Println(speedCount.RenderJSON(roundUpUnit))

	wg.Wait()
}
