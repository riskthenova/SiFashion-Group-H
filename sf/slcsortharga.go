package main

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
