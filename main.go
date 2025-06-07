package main

import "fmt"

type User struct {
	Minat         [100]string
	Keahlian      [100]string
	MinatCount    int
	KeahlianCount int
}

type Karir struct {
	NamaPekerjaan string
	Industri      string
	MatchScore    int
	RataGaji      int
}

var user User
var listKarir = []Karir{
	{"Software Engineer", "Teknologi", 0, 80000},
	{"Data Scientist", "Teknologi", 0, 85000},
	{"Graphic Designer", "Design", 0, 50000},
	{"Manager Akuntasi", "Management", 0, 70000},
	{"Marketing Specialist", "Bisnis", 0, 60000},
	{"Frontend", "Teknologi", 0, 80000},
	{"Backend", "Teknologi", 0, 90000},
}

func main() {
	var choice int
	for {
		fmt.Println("\nCareer Path Recommender")
		fmt.Println("1. Tambah minat/keahlian")
		fmt.Println("2. Ubah minat/keahlian")
		fmt.Println("3. Hapus minat/keahlian")
		fmt.Println("4. Rekomendasi jalur karier")
		fmt.Println("5. Cari jalur karier")
		fmt.Println("6. Urutkan jalur karier")
		fmt.Println("7. Tampilkan statistik")
		fmt.Println("8. Lihat minat dan keahlian")
		fmt.Println("9. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			tambahMinatKeahlian()
		case 2:
			editMinatkehalian()
		case 3:
			deleteMinatkehalian()
		case 4:
			rekomendasiKarir()
		case 5:
			searchKarir()
		case 6:
			sortKarir()
		case 7:
			tampilkanStatistik()
		case 8:
			lihatMinatKeahlian()
		case 9:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Prodecure untuk menerima input dari pengguna lalu menambahkannya ke array minat atau keahlian.
func tambahMinatKeahlian() {
	var t, val string
	fmt.Print("Tambah (minat/keahlian): ")
	fmt.Scan(&t)
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&val)
	if t == "minat" {
		user.Minat[user.MinatCount] = val
		user.MinatCount++
	} else if t == "keahlian" {
		user.Keahlian[user.KeahlianCount] = val
		user.KeahlianCount++
	} else {
		fmt.Println("Tipe tidak valid")
	}
}

// Prodecure untuk mengubah nilai minat atau keahlian yang sudah ada.
// Menggunakan fungsi replaceInArray untuk mengganti nilai lama dengan nilai baru.
func editMinatkehalian() {
	var t, oldVal, newVal string
	fmt.Print("Ubah (minat/keahlian): ")
	fmt.Scan(&t)
	fmt.Print("Nilai lama: ")
	fmt.Scan(&oldVal)
	fmt.Print("Nilai baru: ")
	fmt.Scan(&newVal)
	if t == "minat" {
		replaceInArray(&user.Minat, user.MinatCount, oldVal, newVal)
	} else if t == "keahlian" {
		replaceInArray(&user.Keahlian, user.KeahlianCount, oldVal, newVal)
	} else {
		fmt.Println("Tipe tidak valid")
	}
}

// Prodecure untuk menghapus nilai minat atau keahlian yang sudah ada.
// Menggunakan fungsi deleteFromArray untuk menghapus nilai dari array.
func deleteMinatkehalian() {
	var t, val string
	fmt.Print("Hapus (minat/keahlian): ")
	fmt.Scan(&t)
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&val)
	if t == "minat" {
		deleteFromArray(&user.Minat, &user.MinatCount, val)
	} else if t == "keahlian" {
		deleteFromArray(&user.Keahlian, &user.KeahlianCount, val)
	} else {
		fmt.Println("Tipe tidak valid")
	}
}

// Fungsi untuk mengganti nilai dalam array.
// Menerima pointer ke array, jumlah elemen yang valid, nilai lama, dan nilai baru.
func replaceInArray(arr *[100]string, count int, oldVal, newVal string) {
	for i := 0; i < count; i++ {
		if (*arr)[i] == oldVal {
			(*arr)[i] = newVal
			return
		}
	}
	fmt.Println("Nilai tidak ditemukan")
}

// Fungsi untuk menghapus nilai dari array.
// Menerima pointer ke array, pointer ke jumlah elemen yang valid, dan nilai yang akan dihapus.
func deleteFromArray(arr *[100]string, count *int, val string) {
	for i := 0; i < *count; i++ {
		if (*arr)[i] == val {
			for j := i; j < *count-1; j++ {
				(*arr)[j] = (*arr)[j+1]
			}
			*count--
			return
		}
	}
	fmt.Println("Nilai tidak ditemukan")
}

// Prodecure untuk merekomendasikan jalur karier berdasarkan minat dan keahlian pengguna.
// Menghitung skor kecocokan untuk setiap jalur karier berdasarkan minat dan keahlian.
func rekomendasiKarir() {
	for i := range listKarir {
		listKarir[i].MatchScore = 0
		for j := 0; j < user.MinatCount; j++ {
			if periksaString(listKarir[i].NamaPekerjaan, user.Minat[j]) || periksaString(listKarir[i].Industri, user.Minat[j]) {
				listKarir[i].MatchScore += 10
			}
		}
		for j := 0; j < user.KeahlianCount; j++ {
			if periksaString(listKarir[i].NamaPekerjaan, user.Keahlian[j]) || periksaString(listKarir[i].Industri, user.Keahlian[j]) {
				listKarir[i].MatchScore += 10
			}
		}
	}
	for _, c := range listKarir {
		fmt.Printf("%s (%s) - Match: %d\n", c.NamaPekerjaan, c.Industri, c.MatchScore)
	}
}

// Fungsi untuk memeriksa apakah string str dimulai dengan substring substr.
// Mengembalikan true jika str dimulai dengan substr, false jika tidak.
func periksaString(str, substr string) bool {
	return len(substr) > 0 && len(str) > 0 && (str == substr || (len(str) >= len(substr) && str[:len(substr)] == substr))
}

// Prodecure untuk mencari jalur karier berdasarkan nama atau kategori.
// Menggunakan metode pencarian sequential atau binary tergantung pilihan pengguna.
func searchKarir() {
	var method, query string
	fmt.Print("Metode (sequential/binary): ")
	fmt.Scan(&method)
	fmt.Print("Cari nama/kategori: ")
	fmt.Scan(&query)
	if method == "sequential" {
		for _, c := range listKarir {
			if c.NamaPekerjaan == query || c.Industri == query {
				fmt.Printf("Ditemukan: %s (%s)\n", c.NamaPekerjaan, c.Industri)
				return
			}
		}
		fmt.Println("Tidak ditemukan")
	} else if method == "binary" {
		insertionSortByNama()
		left, right := 0, 6
		for left <= right {
			mid := (left + right) / 2
			if periksaString(listKarir[mid].NamaPekerjaan, query) || periksaString(listKarir[mid].Industri, query) {
				fmt.Printf("Ditemukan: %s (%s)\n", listKarir[mid].NamaPekerjaan, listKarir[mid].Industri)
				return
			} else if listKarir[mid].NamaPekerjaan < query {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		fmt.Println("Tidak ditemukan")
	}
}

// Prodecure untuk mengurutkan jalur karier berdasarkan metode dan kriteria yang dipilih pengguna.
func sortKarir() {
	var method, criteria string
	fmt.Print("Metode (selection/insertion): ")
	fmt.Scan(&method)
	fmt.Print("Kriteria (match/gaji): ")
	fmt.Scan(&criteria)
	if method == "selection" {
		if criteria == "match" {
			selectionSortByMatchScore()
		} else if criteria == "gaji" {
			selectionSortByGaji()
		}
	} else if method == "insertion" {
		if criteria == "match" {
			insertionSortByMatchScore()
		} else if criteria == "gaji" {
			insertionSortBySalary()
		}
	}
	for _, c := range listKarir {
		fmt.Printf("%s (%s) - Match: %d, Gaji: %d\n", c.NamaPekerjaan, c.Industri, c.MatchScore, c.RataGaji)
	}
}

// Prodecure untuk mengurutkan jalur karier berdasarkan skor kecocokan menggunakan algoritma selection sort.
func selectionSortByMatchScore() {
	for i := 0; i < 7; i++ {
		maxIdx := i
		for j := i + 1; j < 7; j++ {
			if listKarir[j].MatchScore > listKarir[maxIdx].MatchScore {
				maxIdx = j
			}
		}
		listKarir[i], listKarir[maxIdx] = listKarir[maxIdx], listKarir[i]
	}
}

// Prodecure untuk mengurutkan jalur karier berdasarkan skor kecocokan menggunakan algoritma insertion sort.
func insertionSortByMatchScore() {
	for i := 1; i < 7; i++ {
		key := listKarir[i]
		j := i - 1
		for j >= 0 && listKarir[j].MatchScore < key.MatchScore {
			listKarir[j+1] = listKarir[j]
			j--
		}
		listKarir[j+1] = key
	}
}

// Prodecure untuk mengurutkan jalur karier berdasarkan gaji menggunakan algoritma selection sort.
func selectionSortByGaji() {
	for i := 0; i < 7; i++ {
		maxIdx := i
		for j := i + 1; j < 7; j++ {
			if listKarir[j].RataGaji > listKarir[maxIdx].RataGaji {
				maxIdx = j
			}
		}
		listKarir[i], listKarir[maxIdx] = listKarir[maxIdx], listKarir[i]
	}
}

// Prodecure untuk mengurutkan jalur karier berdasarkan skor kecocokan menggunakan algoritma insertion sort.
func insertionSortBySalary() {
	for i := 1; i < 7; i++ {
		key := listKarir[i]
		j := i - 1
		for j >= 0 && listKarir[j].RataGaji < key.RataGaji {
			listKarir[j+1] = listKarir[j]
			j--
		}
		listKarir[j+1] = key
	}
}

// Prodecure untuk mengurutkan jalur karier berdasarkan nama menggunakan algoritma insertion sort.
func insertionSortByNama() {
	for i := 1; i < 7; i++ {
		key := listKarir[i]
		j := i - 1
		for j >= 0 && listKarir[j].NamaPekerjaan > key.NamaPekerjaan {
			listKarir[j+1] = listKarir[j]
			j--
		}
		listKarir[j+1] = key
	}
}

func tampilkanStatistik() {
	totalMatch := 0
	for i := 0; i < 7; i++ {
		totalMatch += listKarir[i].MatchScore
	}
	fmt.Println("Statistik kecocokan terhadap jalur karier:")
	for i := 0; i < 7; i++ {
		percent := 0.0
		if totalMatch > 0 {
			percent = float64(listKarir[i].MatchScore) / float64(totalMatch) * 100
		}
		fmt.Printf("%s: %.2f%%\n", listKarir[i].NamaPekerjaan, percent)
	}
}

func lihatMinatKeahlian() {
	fmt.Println("Minat yang telah diinput:")
	if user.MinatCount == 0 {
		fmt.Println("- (belum ada minat)")
	} else {
		for i := 0; i < user.MinatCount; i++ {
			fmt.Printf("- %s\n", user.Minat[i])
		}
	}

	fmt.Println("\nKeahlian yang telah diinput:")
	if user.KeahlianCount == 0 {
		fmt.Println("- (belum ada keahlian)")
	} else {
		for i := 0; i < user.KeahlianCount; i++ {
			fmt.Printf("- %s\n", user.Keahlian[i])
		}
	}
}
