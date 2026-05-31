/*
 * Program      : Aplikasi Manajemen dan Pemantauan Polusi Udara Lokal
 * Workspace    : fjustincia/tugas-besar-polusi
 * Description  : Implementasi CRUD, Sequential/Binary Search, dan Selection/Insertion Sort di Golang
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DataPolusi struct {
	Lokasi   string
	Tanggal  string // Format YYYY-MM-DD
	AQI      int
	Sumber   string
	Kategori string
}

func main() {
	var database []DataPolusi
	scanner := bufio.NewScanner(os.Stdin)
	var pilihan string

	for {
		fmt.Println("\n=======================================================")
		fmt.Println("  SISTEM MANAJEMEN & PEMANTAUAN POLUSI UDARA LOKAL")
		fmt.Println("=======================================================")
		fmt.Println("1. Tambah Data Polusi")
		fmt.Println("2. Tampilkan Semua Data")
		fmt.Println("3. Ubah Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Cari Data Berdasarkan Lokasi (Sequential Search)")
		fmt.Println("6. Cari Data Berdasarkan Lokasi (Binary Search)")
		fmt.Println("7. Urutkan Data Berdasarkan AQI Tertinggi (Selection Sort)")
		fmt.Println("8. Urutkan Data Berdasarkan Tanggal Terbaru (Insertion Sort)")
		fmt.Println("9. Tampilkan Polusi Tertinggi dalam Periode Tertentu")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		pilihan = strings.TrimSpace(scanner.Text())

		switch pilihan {
		case "1":
			tambahData(&database, scanner)
		case "2":
			tampilkanData(database)
		case "3":
			ubahData(&database, scanner)
		case "4":
			hapusData(&database, scanner)
		case "5":
			fmt.Print("Masukkan nama lokasi yang dicari: ")
			scanner.Scan()
			target := scanner.Text()
			sequentialSearch(database, target)
		case "6":
			fmt.Print("Masukkan nama lokasi yang dicari: ")
			scanner.Scan()
			target := scanner.Text()
			binarySearch(database, target)
		case "7":
			selectionSortAQI(database)
			fmt.Println("\n[INFO] Data berhasil diurutkan berdasarkan AQI tertinggi.")
			tampilkanData(database)
		case "8":
			insertionSortTanggal(database)
			fmt.Println("\n[INFO] Data berhasil diurutkan berdasarkan Tanggal terbaru.")
			tampilkanData(database)
		case "9":
			filterTertinggiPeriode(database, scanner)
		case "0":
			fmt.Println("Keluar dari program. Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Menentukan kategori polusi (Spesifikasi b)
func tentukanKategori(aqi int) string {
	if aqi <= 50 {
		return "Baik"
	} else if aqi <= 100 {
		return "Sedang"
	} else if aqi <= 150 {
		return "Tidak Sehat (Sensitif)"
	} else if aqi <= 200 {
		return "Tidak Sehat"
	} else if aqi <= 300 {
		return "Sangat Tidak Sehat"
	}
	return "Berbahaya"
}

// Menampilkan data dalam format tabel
func tampilkanData(data []DataPolusi) {
	if len(data) == 0 {
		fmt.Println("\n[INFO] Data masih kosong!")
		return
	}
	fmt.Println("\n--------------------------------------------------------------------------------------")
	fmt.Printf("%-5s %-20s %-15s %-10s %-20s %s\n", "No", "Lokasi", "Tanggal", "AQI", "Kategori", "Sumber Polusi")
	fmt.Println("--------------------------------------------------------------------------------------")
	for i, item := range data {
		fmt.Printf("%-5d %-20s %-15s %-10d %-20s %s\n", i+1, item.Lokasi, item.Tanggal, item.AQI, item.Kategori, item.Sumber)
	}
	fmt.Println("--------------------------------------------------------------------------------------")
}

// Spesifikasi a & b: Tambah data dan warning ambang batas
func tambahData(data *[]DataPolusi, scanner *bufio.Scanner) {
	var baru DataPolusi
	fmt.Println("\n--- Tambah Data Polusi ---")

	fmt.Print("Masukkan Lokasi: ")
	scanner.Scan()
	baru.Lokasi = scanner.Text()

	fmt.Print("Masukkan Tanggal (YYYY-MM-DD): ")
	scanner.Scan()
	baru.Tanggal = scanner.Text()

	fmt.Print("Masukkan Nilai AQI: ")
	scanner.Scan()
	aqiStr := scanner.Text()
	baru.AQI, _ = strconv.Atoi(aqiStr)

	fmt.Print("Masukkan Sumber Polusi: ")
	scanner.Scan()
	baru.Sumber = scanner.Text()

	baru.Kategori = tentukanKategori(baru.AQI)
	*data = append(*data, baru)

	// Sistem Peringatan
	if baru.AQI > 100 {
		fmt.Println("\n[WARNING] Tingkat polusi melebihi ambang batas aman!")
		fmt.Printf("Status: %s\n", baru.Kategori)
	}
	fmt.Println("[SUCCESS] Data berhasil ditambahkan.")
}

// Fitur Ubah Data (Spesifikasi a)
func ubahData(data *[]DataPolusi, scanner *bufio.Scanner) {
	tampilkanData(*data)
	if len(*data) == 0 {
		return
	}

	fmt.Print("Pilih nomor data yang ingin diubah: ")
	scanner.Scan()
	indexStr := scanner.Text()
	index, err := strconv.Atoi(indexStr)

	if err == nil && index > 0 && index <= len(*data) {
		idx := index - 1

		fmt.Print("Lokasi Baru: ")
		scanner.Scan()
		(*data)[idx].Lokasi = scanner.Text()

		fmt.Print("Tanggal Baru (YYYY-MM-DD): ")
		scanner.Scan()
		(*data)[idx].Tanggal = scanner.Text()

		fmt.Print("Nilai AQI Baru: ")
		scanner.Scan()
		aqiStr := scanner.Text()
		(*data)[idx].AQI, _ = strconv.Atoi(aqiStr)

		fmt.Print("Sumber Polusi Baru: ")
		scanner.Scan()
		(*data)[idx].Sumber = scanner.Text()

		(*data)[idx].Kategori = tentukanKategori((*data)[idx].AQI)
		fmt.Println("[SUCCESS] Data berhasil diubah.")
	} else {
		fmt.Println("[ERROR] Nomor data tidak valid!")
	}
}

// Fitur Hapus Data (Spesifikasi a)
func hapusData(data *[]DataPolusi, scanner *bufio.Scanner) {
	tampilkanData(*data)
	if len(*data) == 0 {
		return
	}

	fmt.Print("Pilih nomor data yang ingin dihapus: ")
	scanner.Scan()
	indexStr := scanner.Text()
	index, err := strconv.Atoi(indexStr)

	if err == nil && index > 0 && index <= len(*data) {
		idx := index - 1
		// Menghapus elemen dari slice
		*data = append((*data)[:idx], (*data)[idx+1:]...)
		fmt.Println("[SUCCESS] Data berhasil dihapus.")
	} else {
		fmt.Println("[ERROR] Nomor data tidak valid!")
	}
}

// Spesifikasi c: Pencarian (Sequential Search)
func sequentialSearch(data []DataPolusi, target string) {
	found := false
	fmt.Printf("\n--- Hasil Sequential Search untuk '%s' ---\n", target)
	for _, item := range data {
		if strings.EqualFold(item.Lokasi, target) {
			fmt.Printf("Lokasi: %s | Tanggal: %s | AQI: %d (%s)\n", item.Lokasi, item.Tanggal, item.AQI, item.Kategori)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Spesifikasi c: Pencarian (Binary Search)
func binarySearch(data []DataPolusi, target string) {
	// Copy data agar tidak merusak urutan asli saat sorting untuk binary search
	dataCopy := make([]DataPolusi, len(data))
	copy(dataCopy, data)

	// Binary search wajib data diurutkan terlebih dahulu
	sort.Slice(dataCopy, func(i, j int) bool {
		return strings.ToLower(dataCopy[i].Lokasi) < strings.ToLower(dataCopy[j].Lokasi)
	})

	left, right := 0, len(dataCopy)-1
	found := false

	fmt.Printf("\n--- Hasil Binary Search untuk '%s' ---\n", target)
	for left <= right {
		mid := left + (right-left)/2
		midLokasi := strings.ToLower(dataCopy[mid].Lokasi)
		targetLower := strings.ToLower(target)

		if midLokasi == targetLower {
			fmt.Printf("Lokasi: %s | Tanggal: %s | AQI: %d (%s)\n", dataCopy[mid].Lokasi, dataCopy[mid].Tanggal, dataCopy[mid].AQI, dataCopy[mid].Kategori)
			found = true
			break
		}
		if midLokasi < targetLower {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Spesifikasi d: Sorting (Selection Sort by AQI Descending - Tertinggi ke terendah)
func selectionSortAQI(data []DataPolusi) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].AQI > data[maxIdx].AQI {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
}

// Spesifikasi d: Sorting (Insertion Sort by Date Descending - Terbaru ke terlama)
func insertionSortTanggal(data []DataPolusi) {
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].Tanggal < key.Tanggal {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// Spesifikasi e: Tampilkan polusi tertinggi dalam periode tertentu
func filterTertinggiPeriode(data []DataPolusi, scanner *bufio.Scanner) {
	if len(data) == 0 {
		fmt.Println("Data masih kosong!")
		return
	}

	fmt.Print("Masukkan Tanggal Awal (YYYY-MM-DD): ")
	scanner.Scan()
	startDate := scanner.Text()

	fmt.Print("Masukkan Tanggal Akhir (YYYY-MM-DD): ")
	scanner.Scan()
	endDate := scanner.Text()

	var filteredData []DataPolusi
	for _, item := range data {
		if item.Tanggal >= startDate && item.Tanggal <= endDate {
			filteredData = append(filteredData, item)
		}
	}

	if len(filteredData) == 0 {
		fmt.Println("Tidak ada data pada rentang tanggal tersebut.")
		return
	}

	// Gunakan selection sort yang sudah ada untuk mengurutkan hasil filter
	selectionSortAQI(filteredData)

	fmt.Printf("\n--- Daftar Wilayah Polusi Tertinggi (%s s/d %s) ---\n", startDate, endDate)
	tampilkanData(filteredData)
}
