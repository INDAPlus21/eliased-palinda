package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	// var slice [][]uint8 
	for i := 0; i < dy; i++ {
		inner_slice := make([]uint8, dx) //[]uint8{uint8(dx)}
		// slice := make([][]uint8, 0, dy)
		for j := 0; j < dx; j++ {
			inner_slice[i] = uint8(5*j) 
		}
		slice[i] = inner_slice
		// slice = append(slice, sub_slice)
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
