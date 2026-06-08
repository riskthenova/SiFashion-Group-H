package main

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
