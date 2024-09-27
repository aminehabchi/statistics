package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR in args")
		return
	}
	filename := os.Args[1]
	if len(filename) < 5 || filename[len(filename)-4:] != ".txt" {
		fmt.Println("Error in filename")
		return
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if len(content) == 0 {
		fmt.Println("Error : Empty file")
		return
	}
	arr := []int{}
	number := strings.Split(string(content), "\n")
	for i := 0; i < len(number); i++ {
		a, _ := strconv.Atoi(number[i])
		if a != 0 {
			arr = append(arr, a)
		}

	}

	//////////////Median////////////////
	Median := Median(arr)
	//////////////Average////////////////
	Average, b := Average(arr)
	//////////////Variance////////////////
	var V int = Variance(arr, b)
	//////////////Standard Deviation////////////////
	var SV int = StandardDeviation(float64(V))
	//////////////Print////////////////
	fmt.Printf("Average: %d\n", Average)
	fmt.Printf("Median: %d\n", Median)
	fmt.Printf("Variance: %d\n", V)
	fmt.Printf("Standard Deviation: %d\n", SV)
}
func Median(arr []int) int {
	slices.Sort(arr)
	c := float64(0)
	if len(arr)%2 == 1 {
		c = float64(arr[len(arr)/2])
	} else {
		c = float64(arr[len(arr)/2] + arr[(len(arr)/2)-1])
		c /= 2
	}
	return round(c)
}
func Average(arr []int) (int, float64) {
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	b := float64(sum) / float64(len(arr))
	return round(b), b
}
func Variance(arr []int, V float64) int {
	length := len(arr)
	var m float64 = 0
	for i := 0; i < len(arr); i++ {
		m = m + math.Pow((float64(arr[i])-V), 2)
	}
	m /= float64(length)
	return round(m)
}
func StandardDeviation(Variance float64) int {
	n := math.Sqrt(float64(Variance))
	return round(n)
}
func round(n float64) int {
	if int(n*10)%10 >= 5 {
		return int(math.Ceil(n))
	}
	return int(math.Floor(n))
}
