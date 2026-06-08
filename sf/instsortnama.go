package main

import (
	"strings"
)

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
