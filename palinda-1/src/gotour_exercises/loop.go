package main

import "log"
import "math"

func Sqrt(x float64) float64 {
	// z := 1.0
	z := float64(1)
	for {
		log.Println(z, x, z*z - x)
		prev_z := z
		z -= (z*z - x) / (2*z)
		if math.Abs(z - prev_z) < 0.000001 {
			return z
		}
	}
}

func main() {
	log.Println("hello world")
	log.Println(Sqrt(2))
	log.Println(math.Sqrt(2))
}