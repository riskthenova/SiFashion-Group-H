package main

import (
	"strings"
)

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
