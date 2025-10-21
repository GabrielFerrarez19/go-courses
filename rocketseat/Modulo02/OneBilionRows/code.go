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

type Measurements struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	start := time.Now()
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	dados := make(map[string]Measurements)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := dados[location]
		if !ok {
			measurement = Measurements{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++

		}

		dados[location] = measurement
	}

	location := make([]string, 0, len(dados))
	for name := range dados {
		location = append(location, name)
	}

	sort.Strings(location)
	fmt.Printf("{")
	for _, name := range location {
		measurements := dados[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f,", name, measurements.Min, measurements.Sum/float64(measurements.Count), measurements.Max)
	}
	fmt.Printf("}\n")

	fmt.Println(time.Since(start))
}
