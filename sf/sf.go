package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type VarianStok struct {
	Ukuran string
	Warna  string
	Jumlah int
}

type Produk struct {
	ID        int
	Nama      string
	Kategori  string
	Harga     float64
	Varian    [10]VarianStok
	JmlVarian int
	TotalStok int
}

const MAX_PRODUK = 100

var daftarProduk [MAX_PRODUK]Produk
var jumlahProduk int = 0
var nextID int = 1

func Input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	teks, _ := reader.ReadString('\n')
	return strings.TrimSpace(teks)
}

func Integer(prompt string) int {
	var angka int
	fmt.Print(prompt)
	fmt.Scan(&angka)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return angka
}

func Float(prompt string) float64 {
	var angka float64
	fmt.Print(prompt)
	fmt.Scan(&angka)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return angka
}

func Garis() {
	fmt.Println("──────────────────────────────────────────────────────")
}

func Judul(judul string) {
	Garis()
	fmt.Printf("  %s\n", judul)
	Garis()
}

func TotalStok(p Produk) int {
	total := 0
	i := 0
	for i < p.JmlVarian {
		total += p.Varian[i].Jumlah
		i++
	}
	return total
}

func TambahProduk() {
	Judul("TAMBAH PRODUK BARU")

	if jumlahProduk >= MAX_PRODUK {
		fmt.Println("  [!] Kapasitas produk penuh! Maksimal", MAX_PRODUK, "produk.")
		return
	}

	var p Produk
	p.ID = nextID
	nextID++

	p.Nama = Input("  Nama Produk      : ")
	p.Kategori = Input("  Kategori (Jenis Pakaian)        : ")
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
			warna := Input("  Warna            : ")
			jumlah := Integer("  Jumlah Stok      : ")

			p.Varian[jmlVarian] = VarianStok{
				Ukuran: strings.ToUpper(ukuran),
				Warna:  strings.Title(strings.ToLower(warna)),
				Jumlah: jumlah,
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

	daftarProduk[jumlahProduk] = p
	jumlahProduk++

	fmt.Printf("\n  [✓] Produk '%s' (ID: %d) berhasil ditambahkan!\n", p.Nama, p.ID)
}

func DetailProduk(p Produk) {
	fmt.Printf("  ID       : %d\n", p.ID)
	fmt.Printf("  Nama     : %s\n", p.Nama)
	fmt.Printf("  Kategori : %s\n", p.Kategori)
	fmt.Printf("  Harga    : Rp %.0f\n", p.Harga)
	fmt.Printf("  Varian   :\n")

	i := 0
	for i < p.JmlVarian {
		fmt.Printf("    [%d] Ukuran: %-5s | Warna: %-15s | Stok: %d\n",
			i+1, p.Varian[i].Ukuran, p.Varian[i].Warna, p.Varian[i].Jumlah)
		i++
	}
	fmt.Printf("  Total Stok: %d pcs\n", p.TotalStok)
}

func SemuaProduk(arr [MAX_PRODUK]Produk, n int) {
	if n == 0 {
		fmt.Println("  Belum ada produk yang tersimpan.")
		return
	}

	fmt.Printf("  %-5s %-25s %-15s %12s %8s\n",
		"ID", "Nama Produk", "Kategori", "Harga (Rp)", "Stok")
	Garis()

	i := 0
	for i < n {
		fmt.Printf("  %-5d %-25s %-15s %12.0f %8d\n",
			arr[i].ID, arr[i].Nama, arr[i].Kategori,
			arr[i].Harga, arr[i].TotalStok)
		i++
	}
}

func UbahProduk() {
	Judul("UBAH DATA PRODUK")

	if jumlahProduk == 0 {
		fmt.Println("  Belum ada produk.")
		return
	}

	idCari := Integer("  Masukkan ID produk yang akan diubah: ")

	idxProduk := SSearchID(daftarProduk, jumlahProduk, idCari)

	if idxProduk == -1 {
		fmt.Println("  [!] Produk dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n  Data produk saat ini:")
	DetailProduk(daftarProduk[idxProduk])
	fmt.Println("\n  (Kosongkan untuk tidak mengubah, tekan Enter)")

	namaBaru := Input("  Nama Baru         : ")
	kategoriBaru := Input("  Kategori Baru     : ")
	hargaStr := Input("  Harga Baru (Rp)   : ")

	if namaBaru != "" {
		daftarProduk[idxProduk].Nama = namaBaru
	}
	if kategoriBaru != "" {
		daftarProduk[idxProduk].Kategori = kategoriBaru
	}
	if hargaStr != "" {
		var hargaBaru float64
		fmt.Sscanf(hargaStr, "%f", &hargaBaru)
		if hargaBaru > 0 {
			daftarProduk[idxProduk].Harga = hargaBaru
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
				varianBaru[jml] = VarianStok{
					Ukuran: strings.ToUpper(ukuran),
					Warna:  strings.Title(strings.ToLower(warna)),
					Jumlah: jumlah,
				}
				jml++
			}
		}
		daftarProduk[idxProduk].Varian = varianBaru
		daftarProduk[idxProduk].JmlVarian = jml
		daftarProduk[idxProduk].TotalStok = TotalStok(daftarProduk[idxProduk])
	}

	fmt.Printf("\n  [✓] Produk ID %d berhasil diperbarui!\n", idCari)
}

func HapusProduk() {
	Judul("HAPUS PRODUK")

	if jumlahProduk == 0 {
		fmt.Println("  Belum ada produk.")
		return
	}

	idCari := Integer("  Masukkan ID produk yang akan dihapus: ")

	idxProduk := SSearchID(daftarProduk, jumlahProduk, idCari)

	if idxProduk == -1 {
		fmt.Println("  [!] Produk tidak ditemukan.")
		return
	}

	fmt.Println("\n  Produk yang akan dihapus:")
	DetailProduk(daftarProduk[idxProduk])
	konfirmasi := Input("\n  Yakin ingin menghapus? (y/n): ")

	if strings.ToLower(konfirmasi) == "y" {
		i := idxProduk
		for i < jumlahProduk-1 {
			daftarProduk[i] = daftarProduk[i+1]
			i++
		}
		daftarProduk[jumlahProduk-1] = Produk{}
		jumlahProduk--
		fmt.Printf("\n  [✓] Produk ID %d berhasil dihapus!\n", idCari)
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}

func SSearchID(arr [MAX_PRODUK]Produk, n int, idCari int) int {
	hasil := -1
	i := 0
	for i < n && hasil == -1 {
		if arr[i].ID == idCari {
			hasil = i
		}
		i++
	}
	return hasil
}

func SSearchUkuran(arr [MAX_PRODUK]Produk, n int, ukuran string) ([MAX_PRODUK]Produk, int) {
	var hasil [MAX_PRODUK]Produk
	jmlHasil := 0
	ukuranUpper := strings.ToUpper(ukuran)

	i := 0
	for i < n {
		j := 0
		ditemukan := false
		for j < daftarProduk[i].JmlVarian && !ditemukan {
			if arr[i].Varian[j].Ukuran == ukuranUpper {
				ditemukan = true
			}
			j++
		}
		if ditemukan {
			hasil[jmlHasil] = arr[i]
			jmlHasil++
		}
		i++
	}
	return hasil, jmlHasil
}

func SSearchWarna(arr [MAX_PRODUK]Produk, n int, warna string) ([MAX_PRODUK]Produk, int) {
	var hasil [MAX_PRODUK]Produk
	jmlHasil := 0
	warnaTitle := strings.Title(strings.ToLower(warna))

	i := 0
	for i < n {
		j := 0
		ditemukan := false
		for j < arr[i].JmlVarian && !ditemukan {
			if arr[i].Varian[j].Warna == warnaTitle {
				ditemukan = true
			}
			j++
		}
		if ditemukan {
			hasil[jmlHasil] = arr[i]
			jmlHasil++
		}
		i++
	}
	return hasil, jmlHasil
}

func UrutkanID(arr [MAX_PRODUK]Produk, n int) [MAX_PRODUK]Produk {
	hasil := arr
	i := 1
	for i < n {
		kunci := hasil[i]
		j := i - 1
		for j >= 0 && hasil[j].ID > kunci.ID {
			hasil[j+1] = hasil[j]
			j--
		}
		hasil[j+1] = kunci
		i++
	}
	return hasil
}

func BSearchID(arr [MAX_PRODUK]Produk, n int, idCari int) int {
	kiri := 0
	kanan := n - 1
	hasil := -1

	for kiri <= kanan && hasil == -1 {
		tengah := (kiri + kanan) / 2
		if arr[tengah].ID == idCari {
			hasil = tengah
		} else if arr[tengah].ID < idCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return hasil
}

func Cari() {
	Judul("PENCARIAN PRODUK")

	if jumlahProduk == 0 {
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
		arrTerurut := UrutkanID(daftarProduk, jumlahProduk)
		fmt.Println("  [Binary Search] Mencari ID", idCari, "...")

		idxBS := BSearchID(arrTerurut, jumlahProduk, idCari)

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

		hasil, jml := SSearchUkuran(daftarProduk, jumlahProduk, ukuran)

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

		hasil, jml := SSearchWarna(daftarProduk, jumlahProduk, warna)

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

func SLCSortHarga(arr [MAX_PRODUK]Produk, n int, ascending bool) [MAX_PRODUK]Produk {
	hasil := arr
	i := 0
	for i < n-1 {
		idxEkstrem := i
		j := i + 1
		for j < n {
			if ascending {
				if hasil[j].Harga < hasil[idxEkstrem].Harga {
					idxEkstrem = j
				}
			} else {
				if hasil[j].Harga > hasil[idxEkstrem].Harga {
					idxEkstrem = j
				}
			}
			j++
		}
		if idxEkstrem != i {
			hasil[i], hasil[idxEkstrem] = hasil[idxEkstrem], hasil[i]
		}
		i++
	}
	return hasil
}

func INSTSortStok(arr [MAX_PRODUK]Produk, n int, ascending bool) [MAX_PRODUK]Produk {
	hasil := arr
	i := 1
	for i < n {
		kunci := hasil[i]
		j := i - 1
		if ascending {
			for j >= 0 && hasil[j].TotalStok > kunci.TotalStok {
				hasil[j+1] = hasil[j]
				j--
			}
		} else {
			for j >= 0 && hasil[j].TotalStok < kunci.TotalStok {
				hasil[j+1] = hasil[j]
				j--
			}
		}
		hasil[j+1] = kunci
		i++
	}
	return hasil
}

func INSTSortNama(arr [MAX_PRODUK]Produk, n int, ascending bool) [MAX_PRODUK]Produk {
	hasil := arr
	i := 1
	for i < n {
		kunci := hasil[i]
		j := i - 1
		if ascending {
			for j >= 0 && strings.ToLower(hasil[j].Nama) > strings.ToLower(kunci.Nama) {
				hasil[j+1] = hasil[j]
				j--
			}
		} else {
			for j >= 0 && strings.ToLower(hasil[j].Nama) < strings.ToLower(kunci.Nama) {
				hasil[j+1] = hasil[j]
				j--
			}
		}
		hasil[j+1] = kunci
		i++
	}
	return hasil
}

func Urut() {
	Judul("TAMPIL & URUTKAN PRODUK")

	if jumlahProduk == 0 {
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
		arrTerurut = SLCSortHarga(daftarProduk, jumlahProduk, ascending)
		labelKriteria = "Harga (Selection Sort)"
	case 2:
		arrTerurut = INSTSortStok(daftarProduk, jumlahProduk, ascending)
		labelKriteria = "Stok (Insertion Sort)"
	default:
		arrTerurut = INSTSortNama(daftarProduk, jumlahProduk, ascending)
		labelKriteria = "Nama (Insertion Sort)"
	}

	fmt.Printf("\n  Sorted by: %s | Urutan: %s\n", labelKriteria, labelArah)
	Garis()
	SemuaProduk(arrTerurut, jumlahProduk)
}

func Statistik() {
	Judul("STATISTIK INVENTARIS")

	if jumlahProduk == 0 {
		fmt.Println("  Belum ada data produk.")
		return
	}

	idxTerbanyak := 0
	idxTersedikit := 0
	i := 1
	for i < jumlahProduk {
		if daftarProduk[i].TotalStok > daftarProduk[idxTerbanyak].TotalStok {
			idxTerbanyak = i
		}
		if daftarProduk[i].TotalStok < daftarProduk[idxTersedikit].TotalStok {
			idxTersedikit = i
		}
		i++
	}

	fmt.Println("  [STOK TERBANYAK]")
	fmt.Printf("  Nama     : %s\n", daftarProduk[idxTerbanyak].Nama)
	fmt.Printf("  Kategori : %s\n", daftarProduk[idxTerbanyak].Kategori)
	fmt.Printf("  Stok     : %d pcs\n", daftarProduk[idxTerbanyak].TotalStok)

	fmt.Println("\n  [STOK TERSEDIKIT]")
	fmt.Printf("  Nama     : %s\n", daftarProduk[idxTersedikit].Nama)
	fmt.Printf("  Kategori : %s\n", daftarProduk[idxTersedikit].Kategori)
	fmt.Printf("  Stok     : %d pcs\n", daftarProduk[idxTersedikit].TotalStok)

	var kategoriList [MAX_PRODUK]string
	var stokPerKategori [MAX_PRODUK]int
	var produkPerKategori [MAX_PRODUK]int
	jmlKategori := 0

	i = 0
	for i < jumlahProduk {
		kat := daftarProduk[i].Kategori
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
			stokPerKategori[jmlKategori] = daftarProduk[i].TotalStok
			produkPerKategori[jmlKategori] = 1
			jmlKategori++
		} else {
			stokPerKategori[idxKat] += daftarProduk[i].TotalStok
			produkPerKategori[idxKat]++
		}
		i++
	}

	fmt.Println("\n  [STOK PER KATEGORI]")
	fmt.Printf("  %-20s %10s %10s\n", "Kategori", "Produk", "Total Stok")
	Garis()

	k := 0
	for k < jmlKategori {
		fmt.Printf("  %-20s %10d %10d pcs\n",
			kategoriList[k], produkPerKategori[k], stokPerKategori[k])
		k++
	}

	totalStokAll := 0
	j := 0
	for j < jumlahProduk {
		totalStokAll += daftarProduk[j].TotalStok
		j++
	}
	Garis()
	fmt.Printf("  Total Produk     : %d item\n", jumlahProduk)
	fmt.Printf("  Total Stok Global: %d pcs\n", totalStokAll)

	fmt.Println("\n  [DISTRIBUSI UKURAN VARIAN]")
	ukuranStd := [5]string{"S", "M", "L", "XL", "XXL"}
	var stokPerUkuran [5]int

	p := 0
	for p < jumlahProduk {
		v := 0
		for v < daftarProduk[p].JmlVarian {
			u := 0
			for u < 5 {
				if daftarProduk[p].Varian[v].Ukuran == ukuranStd[u] {
					stokPerUkuran[u] += daftarProduk[p].Varian[v].Jumlah
				}
				u++
			}
			v++
		}
		p++
	}

	u := 0
	for u < 5 {
		bar := ""
		b := 0
		barLen := stokPerUkuran[u] / 5
		if barLen > 30 {
			barLen = 30
		}
		for b < barLen {
			bar += "█"
			b++
		}
		fmt.Printf("  %-4s | %-30s %d pcs\n", ukuranStd[u], bar, stokPerUkuran[u])
		u++
	}
}

func DataContoh() {
	daftarProduk[0] = Produk{
		ID:       nextID,
		Nama:     "Kemeja Batik Solo",
		Kategori: "Kemeja",
		Harga:    185000,
		Varian: [10]VarianStok{
			{"M", "Biru", 15},
			{"L", "Biru", 20},
			{"XL", "Cokelat", 10},
			{"M", "Cokelat", 12},
		},
		JmlVarian: 4,
	}
	nextID++

	daftarProduk[1] = Produk{
		ID:       nextID,
		Nama:     "Celana Jeans Slim",
		Kategori: "Celana",
		Harga:    275000,
		Varian: [10]VarianStok{
			{"S", "Hitam", 8},
			{"M", "Hitam", 14},
			{"L", "Biru", 18},
			{"XL", "Biru", 6},
		},
		JmlVarian: 4,
	}
	nextID++

	daftarProduk[2] = Produk{
		ID:       nextID,
		Nama:     "Dress Midi Floral",
		Kategori: "Dress",
		Harga:    320000,
		Varian: [10]VarianStok{
			{"S", "Pink", 22},
			{"M", "Pink", 18},
			{"L", "Putih", 10},
		},
		JmlVarian: 3,
	}
	nextID++

	daftarProduk[3] = Produk{
		ID:       nextID,
		Nama:     "Kaos Polos Premium",
		Kategori: "Kaos",
		Harga:    95000,
		Varian: [10]VarianStok{
			{"S", "Putih", 30},
			{"M", "Putih", 35},
			{"L", "Hitam", 28},
			{"XL", "Hitam", 20},
			{"XXL", "Abu", 15},
		},
		JmlVarian: 5,
	}
	nextID++

	daftarProduk[4] = Produk{
		ID:       nextID,
		Nama:     "Jaket Bomber Canvas",
		Kategori: "Jaket",
		Harga:    450000,
		Varian: [10]VarianStok{
			{"M", "Hijau", 7},
			{"L", "Hijau", 9},
			{"XL", "Hitam", 5},
		},
		JmlVarian: 3,
	}
	nextID++

	k := 0
	for k < 5 {
		daftarProduk[k].TotalStok = TotalStok(daftarProduk[k])
		k++
	}

	jumlahProduk = 5
	fmt.Println("  [✓] Data contoh berhasil dimuat (5 produk).")
}

func MenuUtama() {
	Garis()
	fmt.Println("         SiFashion - Sistem Manajemen Inventaris")
	Garis()
	fmt.Println("  [1] Tambah Produk")
	fmt.Println("  [2] Ubah Produk")
	fmt.Println("  [3] Hapus Produk")
	fmt.Println("  [4] Lihat Semua Produk")
	fmt.Println("  [5] Cari Produk (Search)")
	fmt.Println("  [6] Tampil Terurut (Sort)")
	fmt.Println("  [7] Statistik Inventaris")
	fmt.Println("  [8] Muat Data Contoh")
	fmt.Println("  [0] Keluar")
	Garis()
}

func main() {
	fmt.Println()
	Garis()
	fmt.Println("           Selamat Datang di SiFashion")
	fmt.Println("     Sistem Manajemen Inventaris Produk Fashion")

	jalan := true
	for jalan {
		MenuUtama()
		pilihan := Integer("  Pilih menu: ")
		fmt.Println()

		switch pilihan {
		case 1:
			TambahProduk()
		case 2:
			UbahProduk()
		case 3:
			HapusProduk()
		case 4:
			Judul("DAFTAR SEMUA PRODUK")
			SemuaProduk(daftarProduk, jumlahProduk)
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
