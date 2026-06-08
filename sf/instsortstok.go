package main

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
