package main

import (
	"fmt"
)

func DataContoh() {
	inventory.Daftar[0] = Produk{
		ID:       inventory.NextID,
		Nama:     "Kemeja Batik Solo",
		Kategori: "Kemeja",
		Harga:    185000,
		Varian: [10]VarianStok{
			{"M", "Biru", 15, 2},
			{"L", "Biru", 20, 1},
			{"XL", "Cokelat", 10, 1},
			{"M", "Cokelat", 12, 0},
		},
		JmlVarian: 4,
	}
	inventory.NextID++

	inventory.Daftar[1] = Produk{
		ID:       inventory.NextID,
		Nama:     "Celana Jeans Slim",
		Kategori: "Celana",
		Harga:    275000,
		Varian: [10]VarianStok{
			{"S", "Hitam", 8, 2},
			{"M", "Hitam", 14, 3},
			{"L", "Biru", 18, 0},
			{"XL", "Biru", 6, 1},
		},
		JmlVarian: 4,
	}
	inventory.NextID++

	inventory.Daftar[2] = Produk{
		ID:       inventory.NextID,
		Nama:     "Dress Midi Floral",
		Kategori: "Dress",
		Harga:    320000,
		Varian: [10]VarianStok{
			{"S", "Pink", 22, 0},
			{"M", "Pink", 18, 4},
			{"L", "Putih", 10, 2},
		},
		JmlVarian: 3,
	}
	inventory.NextID++

	inventory.Daftar[3] = Produk{
		ID:       inventory.NextID,
		Nama:     "Kaos Polos Premium",
		Kategori: "Kaos",
		Harga:    95000,
		Varian: [10]VarianStok{
			{"S", "Putih", 30, 3},
			{"M", "Putih", 35, 5},
			{"L", "Hitam", 28, 2},
			{"XL", "Hitam", 20, 1},
			{"XXL", "Abu", 15, 0},
		},
		JmlVarian: 5,
	}
	inventory.NextID++

	inventory.Daftar[4] = Produk{
		ID:       inventory.NextID,
		Nama:     "Jaket Bomber Canvas",
		Kategori: "Jaket",
		Harga:    450000,
		Varian: [10]VarianStok{
			{"M", "Hijau", 7, 2},
			{"L", "Hijau", 9, 1},
			{"XL", "Hitam", 5, 4},
		},
		JmlVarian: 3,
	}
	inventory.NextID++

	k := 0
	for k < 5 {
		inventory.Daftar[k].TotalStok = TotalStok(inventory.Daftar[k])
		inventory.Daftar[k].TotalTerjual = HitungTotalTerjual(inventory.Daftar[k])
		k++
	}

	inventory.Jumlah = 5
	fmt.Println("  [✓] Data contoh berhasil dimuat (5 produk).")
}
