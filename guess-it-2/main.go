package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	max := float64(0)
	numbers := []float64{}

	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input!!!")
			return
		}
		if math.Abs(num-max) < 500 {
			max = math.Max(max, num)
			numbers = append(numbers, num)
		}
		avg := Average(numbers)
		m := Median(numbers)
		avg += m
		avg /= 2
		dev := Deviation(numbers)
		ran1 := 0.0
		ran2 := 0.0
		a, b, c := calcul(numbers)
		c *= 10
		if c < 5 {
			ran1 = avg - dev
			ran2 = avg + dev
		} else {
			ran1 = float64(len(numbers)-1)*a + b - 20
			ran2 = float64(len(numbers)-1)*a + b + 20
		}

		fmt.Printf("%f %f\n", math.Abs(ran1), math.Abs(ran2))
	}
}

func Median(nbr []float64) float64 {
	sort.Float64s(nbr)
	l := len(nbr) / 2
	if len(nbr)%2 == 1 {
		return nbr[l]
	}
	return (nbr[l] + nbr[l-1]) / 2
}

func Variance(nbr []float64) float64 {
	avg := Average(nbr)
	var s float64
	for _, v := range nbr {
		s += (v - avg) * (v - avg)
	}
	return s / float64(len(nbr))
}

func Deviation(nbr []float64) float64 {
	return math.Sqrt(Variance(nbr))
}

func Average(nbr []float64) float64 {
	var sum float64
	for _, v := range nbr {
		sum += v
	}
	return sum / float64(len(nbr))
}

func calcul(y []float64) (float64, float64, float64) {
	x := []float64{}
	for i := 0; i < len(y); i++ {
		x = append(x, float64(i))
	}
	var sumxy, sumx, sumy, sumx2, sumy2 float64
	l := float64(len(y))
	for i := 0; i < len(x); i++ {
		sumx2 += (x[i] * x[i])
		sumy2 += (y[i] * y[i])
		sumx += x[i]
		sumy += y[i]
		sumxy += (x[i] * y[i])
	}
	a, b := LinearRegressionLine(sumxy, sumx, sumy, sumx2, sumy2, l)
	c := CorrelationCoefficient(sumxy, sumx, sumy, sumx2, sumy2, l)
	return a, b, c
}

func LinearRegressionLine(sumxy, sumx, sumy, sumx2, sumy2, l float64) (float64, float64) {
	a := float64(l*sumxy-sumx*sumy) / float64(l*sumx2-(sumx*sumx))
	b := float64(sumy)/float64(l) - a*float64(sumx)/float64(l)
	return a, b
}

func CorrelationCoefficient(sumxy, sumx, sumy, sumx2, sumy2, l float64) float64 {
	c := l*sumxy - sumx*sumy
	c /= math.Sqrt((l*sumx2 - sumx*sumx) * (l*sumy2 - sumy*sumy))
	return c
}
