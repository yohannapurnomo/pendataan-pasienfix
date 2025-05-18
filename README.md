# Aplikasi Data Pasien (Golang)

Aplikasi ini dibuat menggunakan bahasa pemrograman Go untuk mencatat dan mengelola data pasien. Fitur yang tersedia meliputi pencatatan data pasien, pengurutan berdasarkan umur dan biaya, serta pencarian pasien berdasarkan nama dan umur.

## 📋 Fitur Utama

1. **Input Data Pasien**  
   Mencatat data pasien berupa nama, umur, penyakit, dan biaya perawatan.

2. **Tampilkan Semua Pasien**  
   Menampilkan daftar seluruh pasien beserta detailnya.

3. **Pengurutan Berdasarkan Umur (Selection Sort)**  
   Mengurutkan daftar pasien dari yang termuda ke yang tertua.

4. **Pengurutan Berdasarkan Biaya (Gnome Sort)**  
   Mengurutkan daftar pasien dari biaya terendah ke tertinggi.

5. **Pencarian Nama (Linear Search)**  
   Mencari pasien berdasarkan nama yang diinput pengguna.

6. **Pencarian Umur (Binary Search)**  
   Mencari pasien berdasarkan umur yang diinput, setelah dilakukan pengurutan.

## 🧠 Struktur Data

Struct `Pasien`:
```go
type Pasien struct {
    Nama     string
    Umur     int
    Penyakit string
    Biaya    int
}
