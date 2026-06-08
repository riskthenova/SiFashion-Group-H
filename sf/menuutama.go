package main

import (
	"fmt"
)

func MenuUtama() {
	Garis()
	fmt.Println("            SiFashion - Sistem Manajemen Inventaris")
	Garis()
	fmt.Println("  [1] Tambah Produk")
	fmt.Println("  [2] Ubah Produk")
	fmt.Println("  [3] Hapus Produk")
	fmt.Println("  [4] Lihat Semua Produk")
	fmt.Println("  [5] Cari Produk")
	fmt.Println("  [6] Tampil Terurut (Sort)")
	fmt.Println("  [7] Statistik Inventaris")
	fmt.Println("  [8] Muat Data Contoh")
	fmt.Println("  [0] Keluar")
	Garis()
}
