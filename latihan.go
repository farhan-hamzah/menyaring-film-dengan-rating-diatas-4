package main

import "fmt"

type Film struct {
	Judul  string
	Rating float64
}

type tabFilm [10]Film

func main() {
	var listFilm tabFilm
	var input Film
	var count, i int
	const maks = 10

	for i < maks && input.Judul != "#" {
		fmt.Scan(&input.Judul)

		if input.Judul != "#" {
			fmt.Scanln(&input.Rating)

			if input.Rating > 4.00 && input.Rating <= 5.00 {
				listFilm[count] = input
				count++
			}
		}
		i++
	}

	// Tampilkan judul film yang memenuhi kriteria
	i = 0
	for i < count {
		fmt.Print(listFilm[i].Judul)
		if i < count-1 {
			fmt.Print(" ")
		}
		i++
	}
	fmt.Println()
}
