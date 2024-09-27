package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	var numbers []float64
	max := float64(0)

	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		if math.Abs(num-max) < 1000 {
			if num > max {
				max = num
			}
			numbers = append(numbers, num)
		}
		avg := Average(numbers)
		dev := Deviation(numbers)
		med := Median(numbers)
		avg += med
		avg /= 2
		fmt.Printf("%f %f\n", math.Abs(avg-dev), avg+dev)
	}
}

func Average(pop []float64) float64 {
	var sum float64
	for _, v := range pop {
		sum += v
	}
	return sum / float64(len(pop))
}

func Median(pop []float64) float64 {
	sort.Float64s(pop)
	l := len(pop) / 2
	if len(pop)%2 == 1 {
		return pop[l]
	}
	return (pop[l] + pop[l-1]) / 2
}

func Variance(pop []float64) float64 {
	avg := Average(pop)
	var s float64
	for _, v := range pop {
		s += (v - avg) * (v - avg)
	}
	return s / float64(len(pop))
}

func Deviation(pop []float64) float64 {
	return math.Sqrt(Variance(pop))
}