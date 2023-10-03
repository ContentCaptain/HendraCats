package main

import "fmt"

// Fungsi ini cek kucing-kucing apakah dewasa atau kitten.
func checkCats(CatsTuti, CatsNining []int) {
	// Buat salinan CatsTuti tanpa yang pertama dan dua terakhir.
	catsTutiCopy := append([]int{}, CatsTuti[1:len(CatsTuti)-2]...)

	// Gabungkan data Tuti yang sudah diperbaiki dengan data Nining.
	allCats := append(catsTutiCopy, CatsNining...)

	// Loop untuk ngecek setiap kucing, dewasa atau kitten, dan tampilin di layar.
	for i, age := range allCats {
		catType := "Kucing Dewasa"
		if age < 3 {
			catType = "Kucing Kecil (Kitten)"
		}
		fmt.Printf("Kucing Nomor %d itu %s, dan umurnya %d tahun\n", i+1, catType, age)
	}
}

func main() {
	// Data uji 1
	dataTuti1 := []int{3, 5, 2, 12, 7}
	dataNining1 := []int{4, 1, 15, 8, 3}
	fmt.Println("Tes Pertama:")
	checkCats(dataTuti1, dataNining1)

	// Data uji 2
	dataTuti2 := []int{9, 16, 6, 8, 3}
	dataNining2 := []int{10, 5, 6, 1, 4}
	fmt.Println("\nTes Kedua:")
	checkCats(dataTuti2, dataNining2)
}
