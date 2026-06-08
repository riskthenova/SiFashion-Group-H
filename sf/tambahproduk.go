package main

import (
	"fmt"
	"strings"
)

func TambahProduk() {
	Judul("TAMBAH PRODUK BARU")

	if inventory.Jumlah >= MAX_PRODUK {
		fmt.Println("  [!] Kapasitas produk penuh! Maksimal", MAX_PRODUK, "produk.")
		return
	}

	var p Produk
	p.ID = inventory.NextID
	inventory.NextID++

	p.Nama = Input("  Nama Produk      : ")
	p.Kategori = Input("  Kategori (Jenis) : ")
	p.Harga = Float("  Harga (Rp)       : ")

	fmt.Println("\n  -- Input Varian Stok (Ukuran & Warna) --")
	fmt.Println("  Ketik 'selesai' pada ukuran untuk berhenti.")

	jmlVarian := 0
	for jmlVarian < 10 {
		fmt.Printf("\n  [Varian %d]\n", jmlVarian+1)
		ukuran := Input("  Ukuran (S/M/L/XL/XXL/selesai): ")

		if strings.ToLower(ukuran) == "selesai" {
			jmlVarian = 10
		} else {
			warna := Input("  Warna                        : ")
			jumlah := Integer("  Jumlah Stok                  : ")
			terjual := Integer("  Jumlah Terjual               : ")

			p.Varian[jmlVarian] = VarianStok{
				Ukuran:  strings.ToUpper(ukuran),
				Warna:   strings.Title(strings.ToLower(warna)),
				Jumlah:  jumlah,
				Terjual: terjual,
			}
			jmlVarian++
			if jmlVarian >= 10 {
				fmt.Println("  [!] Maksimal 10 varian per produk.")
			}
		}
	}

	varian_valid := 0
	k := 0
	for k < 10 {
		if p.Varian[k].Ukuran != "" {
			varian_valid++
		}
		k++
	}
	p.JmlVarian = varian_valid
	p.TotalStok = TotalStok(p)
	p.TotalTerjual = HitungTotalTerjual(p)

	inventory.Daftar[inventory.Jumlah] = p
	inventory.Jumlah++

	fmt.Printf("\n  [✓] Produk '%s' (ID: %d) berhasil ditambahkan!\n", p.Nama, p.ID)
}
