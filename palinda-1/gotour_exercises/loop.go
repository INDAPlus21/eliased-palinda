package main

import "log"
import "math"

func Sqrt(x float64) float64 {
	z := float64(1)
	for {
		prev_z := z
		z -= (z*z - x) / (2*z)
		if math.Abs(z - prev_z) < 0.000001 {
			return z
		}
	}
}

func main() {
	log.Println(Sqrt(2))
	log.Println(math.Sqrt(2))
}