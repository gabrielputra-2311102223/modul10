package main

import (
	"fmt"
)

// Tipe data untuk menyimpan berat balita
type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Subprogram untuk menghitung rata-rata berat balita
func rataRata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBerat arrBalita
	var bMin, bMax float64

	// Input jumlah balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Input berat tiap balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Hitung berat minimum, maksimum, dan rata-rata
	hitungMinMax(arrBerat, n, &bMin, &bMax)
	rata := rataRata(arrBerat, n)

	// Output hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
