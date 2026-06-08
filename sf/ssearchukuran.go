package main

import (
	"strings"
)

func SSearchUkuran(arr [MAX_PRODUK]Produk, n int, ukuran string) ([MAX_PRODUK]Produk, int) {
	var hasil [MAX_PRODUK]Produk
	jmlHasil := 0
	ukuranUpper := strings.ToUpper(ukuran)

	i := 0
	for i < n {
		j := 0
		ditemukan := false
		for j < arr[i].JmlVarian && !ditemukan {
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
