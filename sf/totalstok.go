package main

func TotalStok(p Produk) int {
	total := 0
	i := 0
	for i < p.JmlVarian {
		total += p.Varian[i].Jumlah
		i++
	}
	return total
}

func HitungTotalTerjual(p Produk) int {
	total := 0
	i := 0
	for i < p.JmlVarian {
		total += p.Varian[i].Terjual
		i++
	}
	return total
}
