package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error in args")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	splited := strings.Split(string(data), "\n")
	y := []float64{}
	x := []float64{}
	var j float64 = 0
	for i := 0; i < len(splited); i++ {
		n, err := strconv.ParseFloat(splited[i], 64)
		if err == nil {
			y = append(y, n)
			x = append(x, j)
			j++
		}
	}
	if len(x) == 0 {
		fmt.Println("error!!")
		return
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
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	c := CorrelationCoefficient(sumxy, sumx, sumy, sumx2, sumy2, l)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", c)
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
