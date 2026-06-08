package main

import (
	"fmt"
)

type VarianStok struct {
	Ukuran  string
	Warna   string
	Jumlah  int
	Terjual int
}

type Produk struct {
	ID           int
	Nama         string
	Kategori     string
	Harga        float64
	Varian       [10]VarianStok
	JmlVarian    int
	TotalStok    int
	TotalTerjual int
}

type Inventory struct {
	Daftar [MAX_PRODUK]Produk
	Jumlah int
	NextID int
}

const MAX_PRODUK = 100

var inventory = Inventory{
	Jumlah: 0,
	NextID: 1,
}

func main() {
	fmt.Println()
	Garis()
	fmt.Println("                   Selamat Datang di SiFashion")
	fmt.Println("           Sistem Manajemen Inventaris Produk Fashion")

	jalan := true
	for jalan {
		MenuUtama()
		pilihan := Integer("  Pilih menu: ")

		switch pilihan {
		case 1:
			TambahProduk()
		case 2:
			UbahProduk()
		case 3:
			HapusProduk()
		case 4:
			Judul("DAFTAR SEMUA PRODUK")
			SemuaProduk(inventory.Daftar, inventory.Jumlah)
		case 5:
			Cari()
		case 6:
			Urut()
		case 7:
			Judul("STATISTIK INVENTARIS")
			Statistik()
		case 8:
			Judul("MUAT DATA CONTOH")
			DataContoh()
		case 0:
			fmt.Println("  Terima kasih telah menggunakan SiFashion.")
			fmt.Println("  Program selesai.")
			jalan = false
		default:
			fmt.Println("  [!] Pilihan tidak valid. Silakan coba lagi.")
		}

		if jalan {
			fmt.Println()
			Input("  Tekan Enter untuk kembali ke menu...")
			fmt.Println()
		}
	}
}
