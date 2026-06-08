package main

import (
	"fmt"
)

func Urut() {
	Judul("TAMPIL & URUTKAN PRODUK")

	if inventory.Jumlah == 0 {
		fmt.Println("  Belum ada produk.")
		return
	}

	fmt.Println("  Urutkan berdasarkan:")
	fmt.Println("  [1] Harga    (Selection Sort)")
	fmt.Println("  [2] Stok     (Insertion Sort)")
	fmt.Println("  [3] Nama     (Insertion Sort)")

	pilihan := Integer("\n  Pilih kriteria: ")

	if pilihan < 1 || pilihan > 3 {
		fmt.Println("  [!] Pilihan tidak valid.")
		return
	}

	fmt.Println("\n  Urutan:")
	fmt.Println("  [1] Ascending  (Kecil → Besar / A → Z)")
	fmt.Println("  [2] Descending (Besar → Kecil / Z → A)")

	arah := Integer("\n  Pilih urutan: ")
	ascending := arah == 1

	var arrTerurut [MAX_PRODUK]Produk
	labelKriteria := ""
	labelArah := ""

	if ascending {
		labelArah = "Ascending"
	} else {
		labelArah = "Descending"
	}

	switch pilihan {
	case 1:
		arrTerurut = SLCSortHarga(inventory.Daftar, inventory.Jumlah, ascending)
		labelKriteria = "Harga (Selection Sort)"
	case 2:
		arrTerurut = INSTSortStok(inventory.Daftar, inventory.Jumlah, ascending)
		labelKriteria = "Stok (Insertion Sort)"
	default:
		arrTerurut = INSTSortNama(inventory.Daftar, inventory.Jumlah, ascending)
		labelKriteria = "Nama (Insertion Sort)"
	}

	fmt.Printf("\n  Sorted by: %s | Urutan: %s\n", labelKriteria, labelArah)
	Garis()
	SemuaProduk(arrTerurut, inventory.Jumlah)
}
