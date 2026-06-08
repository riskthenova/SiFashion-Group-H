package main

import (
	"fmt"
)

func Statistik() {
	if inventory.Jumlah == 0 {
		fmt.Println("  Belum ada data produk.")
		return
	}

	idxTerbanyak := 0
	idxTersedikit := 0
	idxTerjualTerbanyak := 0
	idxTerjualTersedikit := 0
	i := 1
	for i < inventory.Jumlah {
		if inventory.Daftar[i].TotalStok > inventory.Daftar[idxTerbanyak].TotalStok {
			idxTerbanyak = i
		}
		if inventory.Daftar[i].TotalStok < inventory.Daftar[idxTersedikit].TotalStok {
			idxTersedikit = i
		}
		if inventory.Daftar[i].TotalTerjual > inventory.Daftar[idxTerjualTerbanyak].TotalTerjual {
			idxTerjualTerbanyak = i
		}
		if inventory.Daftar[i].TotalTerjual < inventory.Daftar[idxTerjualTersedikit].TotalTerjual {
			idxTerjualTersedikit = i
		}
		i++
	}

	fmt.Println("  [STOK TERBANYAK]")
	fmt.Printf("  Nama     : %s\n", inventory.Daftar[idxTerbanyak].Nama)
	fmt.Printf("  Kategori : %s\n", inventory.Daftar[idxTerbanyak].Kategori)
	fmt.Printf("  Stok     : %d pcs\n", inventory.Daftar[idxTerbanyak].TotalStok)

	fmt.Println("\n  [STOK TERSEDIKIT]")
	fmt.Printf("  Nama     : %s\n", inventory.Daftar[idxTersedikit].Nama)
	fmt.Printf("  Kategori : %s\n", inventory.Daftar[idxTersedikit].Kategori)
	fmt.Printf("  Stok     : %d pcs\n", inventory.Daftar[idxTersedikit].TotalStok)

	fmt.Println("\n  [PRODUK TERLARIS]")
	fmt.Printf("  Nama     : %s\n", inventory.Daftar[idxTerjualTerbanyak].Nama)
	fmt.Printf("  Kategori : %s\n", inventory.Daftar[idxTerjualTerbanyak].Kategori)
	fmt.Printf("  Terjual  : %d pcs\n", inventory.Daftar[idxTerjualTerbanyak].TotalTerjual)

	fmt.Println("\n  [PRODUK TERKURANG TERJUAL]")
	fmt.Printf("  Nama     : %s\n", inventory.Daftar[idxTerjualTersedikit].Nama)
	fmt.Printf("  Kategori : %s\n", inventory.Daftar[idxTerjualTersedikit].Kategori)
	fmt.Printf("  Terjual  : %d pcs\n", inventory.Daftar[idxTerjualTersedikit].TotalTerjual)

	var kategoriList [MAX_PRODUK]string
	var stokPerKategori [MAX_PRODUK]int
	var produkPerKategori [MAX_PRODUK]int
	jmlKategori := 0

	i = 0
	for i < inventory.Jumlah {
		kat := inventory.Daftar[i].Kategori
		idxKat := -1
		k := 0
		for k < jmlKategori && idxKat == -1 {
			if kategoriList[k] == kat {
				idxKat = k
			}
			k++
		}
		if idxKat == -1 {
			kategoriList[jmlKategori] = kat
			stokPerKategori[jmlKategori] = inventory.Daftar[i].TotalStok
			produkPerKategori[jmlKategori] = 1
			jmlKategori++
		} else {
			stokPerKategori[idxKat] += inventory.Daftar[i].TotalStok
			produkPerKategori[idxKat]++
		}
		i++
	}

	fmt.Println("\n  [STOK PER KATEGORI]")
	fmt.Printf("  %-20s %10s %14.5s\n", "Kategori", "Produk", "Total")
	Garis()

	k := 0
	for k < jmlKategori {
		fmt.Printf("  %-20s %10d %10d pcs\n",
			kategoriList[k], produkPerKategori[k], stokPerKategori[k])
		k++
	}

	totalStokAll := 0
	j := 0
	for j < inventory.Jumlah {
		totalStokAll += inventory.Daftar[j].TotalStok
		j++
	}
	Garis()
	fmt.Printf("  Total Produk     : %d item\n", inventory.Jumlah)
	fmt.Printf("  Total Stok Global: %d pcs\n", totalStokAll)
}
