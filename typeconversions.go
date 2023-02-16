package typeconversions

import (
	"fmt"
	"strconv"
	"math"
 )

 func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

func stringToInt(str string) int {
	// convert to int
	x, _ := strconv.ParseInt(str, 10, 64)
	return int(x)
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision
	x, _ := strconv.ParseFloat(str, 64)
	return float64(toFixed(x,4))
	
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision

	s := fmt.Sprintf("%.2f", value)

	return s
}

func TestStringConverstions() {
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
	}
	
	if stringToFloat("123.33333333333") != 123.3333 {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}
	
	if FloatToString(float64(1)/float64(3)) != "0.33" {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}