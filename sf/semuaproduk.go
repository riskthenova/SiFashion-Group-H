package main

import (
	"fmt"
)

func SemuaProduk(arr [MAX_PRODUK]Produk, n int) {
	if n == 0 {
		fmt.Println("  Belum ada produk yang tersimpan.")
		return
	}

	fmt.Printf("  %-5s %-22s %-10s %-10s %-5s %-6s\n",
		"ID", "Nama Produk", "Kategori", "Harga(Rp)", "Stok", "Terjual")
	Garis()

	i := 0
	for i < n {
		fmt.Printf("  %-5d %-22s %-10s %-10.0f %-5d %-6d\n",
			arr[i].ID, arr[i].Nama, arr[i].Kategori,
			arr[i].Harga, arr[i].TotalStok, arr[i].TotalTerjual)
		i++
	}
}
