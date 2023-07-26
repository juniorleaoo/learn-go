package average

import (
	"fmt"
	"strconv"
)

func Average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	avg := total / float64(len(numbers))
	rounded, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
	return rounded
}
