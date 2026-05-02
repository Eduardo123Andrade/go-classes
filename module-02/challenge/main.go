package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Measurement struct {
	Nome  string
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurement, err := os.Open("measurements.txt")

	if err != nil {
		panic(err)
	}
	defer measurement.Close()

	scanner := bufio.NewScanner(measurement)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")
		location := rawData[:semicolonIndex]
		temperature := rawData[semicolonIndex+1:]

		fmt.Println(location, temperature)

	}
}
