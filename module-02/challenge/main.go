package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
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

	dados := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurement)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolonIndex := strings.Index(rawData, ";")
		location := rawData[:semicolonIndex]
		rawTemperature := rawData[semicolonIndex+1:]

		temperature, _ := strconv.ParseFloat(rawTemperature, 64)
		measurement, ok := dados[location]

		if !ok {
			measurement = Measurement{
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temperature)
			measurement.Max = max(measurement.Min, temperature)
			measurement.Sum += temperature
			measurement.Count++
		}

		dados[location] = measurement
	}

	for name, measmeasurement := range dados {
		fmt.Printf("%s: %#+v\n", name, measmeasurement)
	}
}
