package main

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
