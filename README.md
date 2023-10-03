# Pengecekan Umur Kucing

Program Go ini digunakan untuk menentukan apakah kucing-kucing sudah dewasa atau masih anak kucing berdasarkan umur mereka. Program ini menerima dua rangkaian data berupa angka integer, yaitu CatsTuti dan CatsNining, yang mewakili umur kucing-kucing milik dua individu, Tuti dan Nining, dan menentukan status umur mereka.

## Daftar Isi
1. [Instalasi](#instalasi)
2. [Penggunaan](#penggunaan)
3. [Contoh](#contoh)
4. [Berkontribusi](#berkontribusi)
5. [Lisensi](#lisensi)

## Instalasi
Untuk menggunakan program ini, Anda perlu memiliki Go yang terinstal di komputer Anda. Jika belum, Anda dapat mengunduh dan menginstalnya dari [situs web resmi Go](https://golang.org/).

Setelah Go terinstal, Anda dapat mengklon repositori ini ke komputer lokal Anda dengan perintah berikut:

```bash
git clone https://github.com/ContentCaptain/HendraCats


#Penggunaan
Untuk menggunakan fungsi checkCats untuk menentukan apakah kucing-kucing sudah dewasa atau masih anak kucing, Anda perlu memanggilnya dengan dua rangkaian angka integer yang mewakili umur kucing. Berikut cara menggunakannya:
import "fmt"

// ...

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

#Contoh
Berikut adalah beberapa contoh bagaimana program ini berfungsi:

Tes Pertama:
Kucing Nomor 1 itu Kucing Dewasa, dan umurnya 5 tahun
Kucing Nomor 2 itu Kucing Dewasa, dan umurnya 2 tahun
Kucing Nomor 3 itu Kucing Kecil (Kitten), dan umurnya 12 tahun
Kucing Nomor 4 itu Kucing Kecil (Kitten), dan umurnya 7 tahun

Tes Kedua:
Kucing Nomor 1 itu Kucing Dewasa, dan umurnya 16 tahun
Kucing Nomor 2 itu Kucing Dewasa, dan umurnya 6 tahun
Kucing Nomor 3 itu Kucing Dewasa, dan umurnya 8 tahun
Kucing Nomor 4 itu Kucing Kecil (Kitten), dan umurnya 3 tahun
Kucing Nomor 5 itu Kucing Kecil (Kitten), dan umurnya 10 tahun


#Berkontribusi
Jika Anda ingin berkontribusi ke proyek ini atau melaporkan masalah, silakan buka isu atau kirim permintaan tarik di repositori GitHub.

#Lisensi
Program ini bersifat open-source dan tersedia di bawah Lisensi MIT.
