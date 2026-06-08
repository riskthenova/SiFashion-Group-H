package main

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
