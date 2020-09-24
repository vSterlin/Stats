package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/montanaflynn/stats"
)

func main() {
	jsonFile, err := os.Open(os.Args[1] + ".json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	arr := []float64{}
	json.Unmarshal(byteValue, &arr)

	sort.Float64s(arr)

	v, _ := stats.Variance(arr)
	sd, _ := stats.StandardDeviation(arr)

	mean, _ := stats.Mean(arr)
	median, _ := stats.Median(arr)
	mode, _ := stats.Mode(arr)
	fmt.Printf("Variance: %0.4f\n", v)
	fmt.Printf("SD: %0.4f\n", sd)
	fmt.Printf("Range: %0.4f\n", arr[len(arr)-1]-arr[0])

	fmt.Printf("Mean: %0.4f\nMedian: %0.4f\nMode: %0.4f\n", mean, median, mode)

	quartiles, _ := stats.Quartile(arr)
	fmt.Printf("Q1: %0.4f, Q2: %0.4f, Q3: %0.4f\n", quartiles.Q1, quartiles.Q2, quartiles.Q3)
	iqr, _ := stats.InterQuartileRange(arr)
	fmt.Printf("IQR: %0.4f\n", iqr)

}
