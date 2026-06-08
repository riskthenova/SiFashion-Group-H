package main

import (
	"fmt"
)

func DetailProduk(p Produk) {
	fmt.Printf("  ID       : %d\n", p.ID)
	fmt.Printf("  Nama     : %s\n", p.Nama)
	fmt.Printf("  Kategori : %s\n", p.Kategori)
	fmt.Printf("  Harga    : Rp %.0f\n", p.Harga)
	fmt.Printf("  Varian   :\n")

	i := 0
	for i < p.JmlVarian {
		fmt.Printf("  [%d] Ukuran: %-5s | Warna: %-5s | Stok: %d | Terjual: %d\n",
			i+1, p.Varian[i].Ukuran, p.Varian[i].Warna, p.Varian[i].Jumlah, p.Varian[i].Terjual)
		i++
	}
	fmt.Printf("  Total Stok: %d pcs\n", p.TotalStok)
	fmt.Printf("  Total Terjual: %d pcs\n", p.TotalTerjual)
}
