package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	start := time.Now()
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

	locations := make([]string, 0, len(dados))

	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Print("{")
	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Println("}\n")
	fmt.Println(time.Since(start))
}
