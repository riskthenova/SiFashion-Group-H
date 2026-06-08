package main

import (
	"fmt"
	"strings"
)

func HapusProduk() {
	Judul("HAPUS PRODUK")

	if inventory.Jumlah == 0 {
		fmt.Println("  Belum ada produk.")
		return
	}

	idCari := Integer("  Masukkan ID produk yang akan dihapus: ")

	idxProduk := SSearchID(inventory.Daftar, inventory.Jumlah, idCari)

	if idxProduk == -1 {
		fmt.Println("  [!] Produk tidak ditemukan.")
		return
	}

	fmt.Println("\n  Produk yang akan dihapus:")
	DetailProduk(inventory.Daftar[idxProduk])
	konfirmasi := Input("\n  Yakin ingin menghapus? (y/n): ")

	if strings.ToLower(konfirmasi) == "y" {
		i := idxProduk
		for i < inventory.Jumlah-1 {
			inventory.Daftar[i] = inventory.Daftar[i+1]
			i++
		}
		inventory.Daftar[inventory.Jumlah-1] = Produk{}
		inventory.Jumlah--
		fmt.Printf("\n  [✓] Produk ID %d berhasil dihapus!\n", idCari)
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}
