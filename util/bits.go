package util

import "math"

// Returns the number of bits needed to store the specified number of values
func Bits(values int) uint {
	if values <= 0 {
		return 0
	} else if values == 1 {
		return 1
	} else {
		return uint(math.Ceil(math.Log2(float64(values))))
	}
}
