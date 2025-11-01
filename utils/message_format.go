package utils

import (
	"fmt"
	"strconv"
)

func FormatMs(ms float64) string {
	return fmt.Sprintf("%.2f ms", ms)
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}
