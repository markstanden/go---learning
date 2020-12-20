package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// take the rawReaction data from the given text file, convert from bytes to string
	rawReactions := string(bsFromFile("./reaction_times.txt"))

	// split the string into a slice, splitting at the end of line characters
	reactions := strings.Split(rawReactions, "\r\n")

	//declare an empty sum float variable
	var sum float64

	//iterate the slice, adding the floats to the declared float
	for _, i := range reactions {
		sum += stringToFloat(i)
	}

	//divide the float sum by the length of the slice to get the mean average
	meanFloat := sum / float64(len(reactions))

	//round the number to 1DP
	mean1DP := math.Round(meanFloat*10) / 10

	//output to stdout
	fmt.Println(mean1DP)
}

/* Function to convert string value to float64 value */
func stringToFloat(str string) float64 {
	//parse the string to float, returns flt if successful, returns err if it fails
	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Float parse error:", err)
		// exit on err
		os.Exit(1)
	}
	return flt
}
