package main

import (
	"fmt"
	"strings"
)

func Cari() {
	Judul("PENCARIAN PRODUK")

	if inventory.Jumlah == 0 {
		fmt.Println("  Belum ada produk.")
		return
	}

	fmt.Println("  Cari berdasarkan:")
	fmt.Println("  [1] ID Produk (Binary Search)")
	fmt.Println("  [2] Ukuran    (Sequential Search)")
	fmt.Println("  [3] Warna     (Sequential Search)")

	pilihan := Integer("\n  Pilih metode: ")

	switch pilihan {
	case 1:
		idCari := Integer("  Masukkan ID: ")

		fmt.Println("\n  [Binary Search] Mengurutkan array berdasarkan ID...")
		arrTerurut := UrutkanID(inventory.Daftar, inventory.Jumlah)
		fmt.Println("  [Binary Search] Mencari ID", idCari, "...")

		idxBS := BSearchID(arrTerurut, inventory.Jumlah, idCari)

		if idxBS == -1 {
			fmt.Println("  [!] Produk tidak ditemukan (Binary Search).")
		} else {
			fmt.Println("\n  [✓] Produk ditemukan (Binary Search):")
			Garis()
			DetailProduk(arrTerurut[idxBS])
		}

	case 2:
		ukuran := Input("  Masukkan ukuran (S/M/L/XL/XXL): ")
		fmt.Println("\n  [Sequential Search] Mencari produk dengan ukuran", strings.ToUpper(ukuran), "...")

		hasil, jml := SSearchUkuran(inventory.Daftar, inventory.Jumlah, ukuran)

		if jml == 0 {
			fmt.Println("  [!] Tidak ada produk dengan ukuran tersebut.")
		} else {
			fmt.Printf("\n  [✓] Ditemukan %d produk:\n", jml)
			Garis()
			SemuaProduk(hasil, jml)
		}

	case 3:
		warna := Input("  Masukkan warna: ")
		fmt.Println("\n  [Sequential Search] Mencari produk dengan warna", strings.Title(strings.ToLower(warna)), "...")

		hasil, jml := SSearchWarna(inventory.Daftar, inventory.Jumlah, warna)

		if jml == 0 {
			fmt.Println("  [!] Tidak ada produk dengan warna tersebut.")
		} else {
			fmt.Printf("\n  [✓] Ditemukan %d produk:\n", jml)
			Garis()
			SemuaProduk(hasil, jml)
		}
	default:
		fmt.Println("  [!] Pilihan tidak valid.")
	}
}
