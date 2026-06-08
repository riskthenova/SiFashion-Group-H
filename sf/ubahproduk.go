package main

import (
	"fmt"
	"strings"
)

func UbahProduk() {
	Judul("UBAH DATA PRODUK")

	if inventory.Jumlah == 0 {
		fmt.Println("  Belum ada produk.")
		return
	}

	idCari := Integer("  Masukkan ID produk yang akan diubah: ")

	idxProduk := SSearchID(inventory.Daftar, inventory.Jumlah, idCari)

	if idxProduk == -1 {
		fmt.Println("  [!] Produk dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n  Data produk saat ini:")
	DetailProduk(inventory.Daftar[idxProduk])
	fmt.Println("\n  (Kosongkan untuk tidak mengubah, tekan Enter)")

	namaBaru := Input("  Nama Baru         : ")
	kategoriBaru := Input("  Kategori Baru     : ")
	hargaStr := Input("  Harga Baru (Rp)   : ")

	if namaBaru != "" {
		inventory.Daftar[idxProduk].Nama = namaBaru
	}
	if kategoriBaru != "" {
		inventory.Daftar[idxProduk].Kategori = kategoriBaru
	}
	if hargaStr != "" {
		var hargaBaru float64
		fmt.Sscanf(hargaStr, "%f", &hargaBaru)
		if hargaBaru > 0 {
			inventory.Daftar[idxProduk].Harga = hargaBaru
		}
	}

	fmt.Println("\n  Update stok varian? (y/n)")
	pilihan := Input("  Pilihan: ")
	if strings.ToLower(pilihan) == "y" {
		fmt.Println("  Masukkan ulang semua varian (data lama akan diganti):")
		var varianBaru [10]VarianStok
		jml := 0

		lanjut := true
		for lanjut && jml < 10 {
			fmt.Printf("\n  [Varian %d]\n", jml+1)
			ukuran := Input("  Ukuran (atau 'selesai'): ")
			if strings.ToLower(ukuran) == "selesai" {
				lanjut = false
			} else {
				warna := Input("  Warna: ")
				jumlah := Integer("  Jumlah Stok: ")
				terjual := Integer("  Jumlah Terjual: ")
				varianBaru[jml] = VarianStok{
					Ukuran:  strings.ToUpper(ukuran),
					Warna:   strings.Title(strings.ToLower(warna)),
					Jumlah:  jumlah,
					Terjual: terjual,
				}
				jml++
			}
		}
		inventory.Daftar[idxProduk].Varian = varianBaru
		inventory.Daftar[idxProduk].JmlVarian = jml
		inventory.Daftar[idxProduk].TotalStok = TotalStok(inventory.Daftar[idxProduk])
		inventory.Daftar[idxProduk].TotalTerjual = HitungTotalTerjual(inventory.Daftar[idxProduk])
	}

	fmt.Printf("\n  [✓] Produk ID %d berhasil diperbarui!\n", idCari)
}
