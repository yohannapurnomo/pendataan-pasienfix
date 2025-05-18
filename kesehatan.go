package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

type Pasien struct {
Nama     string
Umur     int
Penyakit string
Biaya    int
}

var daftarPasien []Pasien

func inputPasien() {
reader := bufio.NewReader(os.Stdin)
fmt.Print("Masukkan Nama Pasien: ")
nama, _ := reader.ReadString('\n')
nama = strings.TrimSpace(nama)
fmt.Print("Masukkan Umur Pasien: ")

var umur int

fmt.Scanln(&umur)
fmt.Print("Masukkan Penyakit: ")
penyakit, _ := reader.ReadString('\n')
penyakit = strings.TrimSpace(penyakit)
fmt.Print("Masukkan Biaya: ")

var biaya int

fmt.Scanln(&biaya)
baru := Pasien{Nama: nama, Umur: umur, Penyakit: penyakit, Biaya: biaya}
daftarPasien = append(daftarPasien, baru)
fmt.Println("Pasien berhasil ditambahkan!\n")
}

func tampilkanPasien() {
if len(daftarPasien) == 0 {
	fmt.Println("Belum ada data pasien.\n")
	return
}
fmt.Println("=== Daftar Pasien ===")
for i, p := range daftarPasien {
	fmt.Printf("%d. Nama: %s | Umur: %d | Penyakit: %s | Biaya: Rp%d\n", i+1, p.Nama, p.Umur, p.Penyakit, p.Biaya)
}
fmt.Println()
}

func selectionSortUmur(data []Pasien) {
n := len(data)
for i := 0; i < n-1; i++ {
	minIdx := i
	for j := i + 1; j < n; j++ {
		if data[j].Umur < data[minIdx].Umur {
			minIdx = j
		}
	}
	data[i], data[minIdx] = data[minIdx], data[i]
}
}

func gnomeSortBiaya(data []Pasien) {
index := 0
for index < len(data) {
	if index == 0 {
		index++
	if data[index].Biaya >= data[index-1].Biaya {
		index++
	} else {
		data[index], data[index-1] = data[index-1], data[index]
		index--
	}
	}
}
}

func linearSearchNama(data []Pasien, keyword string) {
found := false
for _, p := range data {
	if strings.EqualFold(p.Nama, keyword) {
		fmt.Println("Pasien ditemukan:")
		fmt.Printf("Nama: %s | Umur: %d | Penyakit: %s | Biaya: Rp%d\n\n", p.Nama, p.Umur, p.Penyakit, p.Biaya)
		found = true
		break
	}
}
if !found {
	fmt.Println("Pasien tidak ditemukan.\n")
}
}

func binarySearchUmur(data []Pasien, target int) {
if len(data) == 0 {
	fmt.Println("Tidak ada data pasien.\n")
	return
}

selectionSortUmur(data)
low := 0
high := len(data) - 1
found := false
for low <= high {
	mid := (low + high) / 2
	if data[mid].Umur == target {
		fmt.Println("Pasien dengan umur ditemukan:")
		fmt.Printf("Nama: %s | Umur: %d | Penyakit: %s | Biaya: Rp%d\n\n", data[mid].Nama, data[mid].Umur, data[mid].Penyakit, data[mid].Biaya)
		found = true
		break
	} else if data[mid].Umur < target {
		low = mid + 1
	} else {
		high = mid - 1
	}
}
if !found {
	fmt.Println("Tidak ada pasien dengan umur tersebut.\n")
}
}

func menu() {
for {
	fmt.Println("====== Menu Utama ======")
	fmt.Println("1. Input Data Pasien")
	fmt.Println("2. Tampilkan Semua Pasien")
	fmt.Println("3. Urutkan Berdasarkan Umur (Selection Sort)")
	fmt.Println("4. Urutkan Berdasarkan Biaya (Gnome Sort)")
	fmt.Println("5. Cari Pasien Berdasarkan Nama (Linear Search)")
	fmt.Println("6. Cari Pasien Berdasarkan Umur (Binary Search)")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu (0-6): ")

	var pilihan string
	fmt.Scanln(&pilihan)
	switch pilihan {
	case "1":
		inputPasien()
	case "2":
		tampilkanPasien()
	case "3":
		selectionSortUmur(daftarPasien)
		fmt.Println("Diurutkan berdasarkan umur!\n")
	case "4":
		gnomeSortBiaya(daftarPasien)
		fmt.Println("Diurutkan berdasarkan biaya!\n")
	case "5":
		fmt.Print("Masukkan nama yang dicari: ")
		reader := bufio.NewReader(os.Stdin)
		nama, _ := reader.ReadString('\n')
		nama = strings.TrimSpace(nama)
		linearSearchNama(daftarPasien, nama)
	case "6":
		fmt.Print("Masukkan umur yang dicari: ")
		var umurStr string
		fmt.Scanln(&umurStr)
		umur, err := strconv.Atoi(umurStr)

		if err == nil {
			binarySearchUmur(daftarPasien, umur)
		} else {
			fmt.Println("Input umur tidak valid.\n")

		}
	case "0":
		fmt.Println("Keluar dari program. Sampai jumpa!")
		return
	default:
		fmt.Println("Pilihan tidak valid. Coba lagi.\n")
	}
}
}

func main() {
menu()
}
