package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		inner_slice := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			inner_slice[i] = uint8(5*j) 
		}
		slice[i] = inner_slice
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
