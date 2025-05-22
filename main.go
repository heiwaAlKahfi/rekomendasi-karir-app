package main

import "fmt"

const NMinat int = 15
const NKeahlian int = 17
const MaxPilihan int = 100
const MaxRekomendasi int = 100

var minatList = []string{
	"Teknologi",
	"Seni",
	"Bisnis",
	"Pendidikan",
	"Kesehatan",
	"Hukum",
	"Lingkungan",
	"Transportasi",
	"Sosial",
	"Ekonomi",
	"Pariwisata",
	"Pangan",
	"Media dan Jurnalistik",
	"Olahraga",
	"Pertanian",
}

var keahlianList = []string{
	"Pemrograman",
	"Desain Grafis",
	"Komunikasi",
	"Manajemen",
	"Analisis Data",
	"Penulisan",
	"Public Speaking",
	"Analisis Hukum",
	"Mekanik",
	"Statistik",
	"Negosiasi",
	"Observasi",
	"Fotografi",
	"Video Editing",
	"Kepemimpinan",
	"Bahasa Asing",
	"Kuliner",
}


var rekomendasiPekerjaan = map[string]map[string]string{
	"Teknologi": {
		"Pemrograman":    "Software Engineer",
		"Analisis Data":  "Data Scientist",
		"Statistik":      "AI Researcher",
		"Manajemen":      "IT Project Manager",
	},
	"Seni": {
		"Desain Grafis":  "Graphic Designer",
		"Penulisan":      "Creative Writer",
		"Video Editing":  "Video Editor",
		"Fotografi":      "Fotografer Artistik",
	},
	"Bisnis": {
		"Manajemen":      "Business Manager",
		"Analisis Data":  "Market Analyst",
		"Negosiasi":      "Sales Executive",
		"Kepemimpinan":   "Startup Founder",
	},
	"Pendidikan": {
		"Public Speaking": "Dosen",
		"Komunikasi":      "Guru",
		"Penulisan":       "Penulis Buku Edukasi",
		"Bahasa Asing":    "Guru Bahasa Asing",
	},
	"Kesehatan": {
		"Observasi":      "Analis Medis",
		"Manajemen":      "Manajer Rumah Sakit",
	},
	"Hukum": {
		"Analisis Hukum": "Pengacara",
		"Public Speaking": "Jaksa",
	},
	"Lingkungan": {
		"Observasi":      "Peneliti Lingkungan",
		"Statistik":      "Analis Dampak Lingkungan",
	},
	"Transportasi": {
		"Mekanik":        "Teknisi Otomotif",
		"Manajemen":      "Manajer Logistik",
	},
	"Sosial": {
		"Komunikasi":     "Pekerja Sosial",
		"Public Speaking": "Motivator",
	},
	"Ekonomi": {
		"Analisis Data":  "Ekonom",
		"Statistik":      "Peneliti Ekonomi",
	},
	"Pariwisata": {
		"Bahasa Asing":   "Pemandu Wisata",
		"Komunikasi":     "Travel Consultant",
	},
	"Pangan": {
		"Kuliner":        "Chef",
		"Manajemen":      "Food Service Manager",
	},
	"Media dan Jurnalistik": {
		"Penulisan":      "Jurnalis",
		"Video Editing":  "Editor Berita",
		"Fotografi":      "Jurnalis Foto",
	},
	"Olahraga": {
		"Kepemimpinan":   "Pelatih Tim",
		"Public Speaking": "Konsultan Kebugaran",
	},
	"Pertanian": {
		"Observasi":      "Agronom",
		"Manajemen":      "Manajer Perkebunan",
	},
}

// Menampikan daftar pilihan berdasarkan parameter yang diberikan (minat atau keahlian).
func tampilkanDaftar(judul string, list []string) {
	fmt.Println("\nPilih " + judul + " (masukkan angka, pisahkan dengan spasi):")
	for i := 0; i < NMinat && judul == "Minat"; i++ {
		fmt.Printf("%d. %s\n", i+1, list[i])
	}
	for i := 0; i < NKeahlian && judul == "Keahlian"; i++ {
		fmt.Printf("%d. %s\n", i+1, list[i])
	}
}

// Mengembalikan nilai array integer dari indeks yang dipilih dan jumlah pilihan yang valid.
// Mengabaikan input yang tidak valid dan menghentikan input jika 0 dimasukkan.
// Jika tidak ada input yang valid, mengembalikan array kosong.
func inputPilihan(placeholder string, jumlah int) ([MaxPilihan]int, int) {
	var temp int
	var hasil [MaxPilihan]int
	idx := 0

	fmt.Println(placeholder)

	for {
		n, err := fmt.Scan(&temp)
		if n == 0 || err != nil || temp == 0 {
			break
		}
		if temp >= 1 && temp <= jumlah {
			hasil[idx] = temp - 1
			idx++
			if idx >= MaxPilihan {
				break
			}
		}
	}
	return hasil, idx
}

// Mengembalikan nilai array string pekerjaan berdasarkan minat dan keahlian yang dipilih dan nilai jumlahHasil dari total rekomendasi pekerjaaan yang cocok.
// Membandingkan array minat dan keahlian yang dipilih dengan rekomendasi pekerjaan.
// Jika tidak ada kombinasi yang cocok, mengembalikan array kosong.
func cariPekerjaan(minatIdxs [MaxPilihan]int, jmMinat int, keahlianIdxs [MaxPilihan]int, jmKeahlian int) ([MaxRekomendasi]string, int) {
	var hasil [MaxRekomendasi]string
	jmHasil := 0
	sudah := map[string]bool{}

	for i := 0; i < jmMinat; i++ {
		for j := 0; j < jmKeahlian; j++ {
			minat := minatList[minatIdxs[i]]
			keahlian := keahlianList[keahlianIdxs[j]]
			if pekerjaan, ok := rekomendasiPekerjaan[minat][keahlian]; ok && !sudah[pekerjaan] {
				hasil[jmHasil] = pekerjaan
				sudah[pekerjaan] = true
				jmHasil++
				if jmHasil >= MaxRekomendasi {
					break
				}
			}
		}
	}
	return hasil, jmHasil
}

func main() {
	var nama string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&nama)

	tampilkanDaftar("Minat", minatList)
	fmt.Println("")
	minatIdxs, jmMinat := inputPilihan("Pilih minat anda (ketik 0 untuk berhenti)", NMinat)

	tampilkanDaftar("Keahlian", keahlianList)
	fmt.Println("")
	keahlianIdxs, jmKeahlian := inputPilihan("Pilih keahlian anda (ketik 0 untuk berhenti)", NKeahlian)

	rekomendasi, jumlahRek := cariPekerjaan(minatIdxs, jmMinat, keahlianIdxs, jmKeahlian)

	fmt.Println("\n====================================")
	fmt.Printf("Halo %s,\nberdasarkan pilihan Anda, berikut rekomendasi karir yang cocok:\n", nama)
	if jumlahRek == 0 {
		fmt.Println("Maaf, belum ada rekomendasi untuk kombinasi tersebut.")
	} else {
		for i := 0; i < jumlahRek; i++ {
			fmt.Printf("- %s\n", rekomendasi[i])
		}
	}
	fmt.Println("====================================")
}
