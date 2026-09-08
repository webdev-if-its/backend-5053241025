package main

import "fmt"

// TODO(Level 1): lihat SOAL.md untuk kontrak lengkap tiap fungsi di bawah.
// Ganti setiap "panic" dengan implementasi yang benar.

func HitungSubtotal(qty int, hargaSatuan float64) float64 {
	var quantity = qty
	var harga = hargaSatuan

	var total int = quantity * int(harga)
	return float64(total)
}

func HitungTotalPesanan(qty []int, hargaSatuan []float64) float64 {
	var total float64 = 0

	if len(qty) != len(hargaSatuan) {
		return 0
	}

	for i := 0; i < len(qty); i++ {
		total += float64(qty[i]) * hargaSatuan[i]
	}
	return total
}

func TerapkanPajak(total float64, tarifPajak float64) float64 {
	var hasil float64

	hasil = total + (total * tarifPajak)
	return hasil
}

func HitungDiskon(total float64) float64 {
	switch {
	case total >= 1000000:
		return total / 10
	case total >= 500000:
		return total * 0.05
	default:
		return 0
	}
}

func TotalSetelahDiskon(qty []int, hargaSatuan []float64, tarifPajak float64) float64 {
	var total float64

	total = HitungTotalPesanan(qty, hargaSatuan)

	total -= HitungDiskon(total)
	return TerapkanPajak(total, tarifPajak)

}

func ValidasiPesanan(qty []int, hargaSatuan []float64) (bool, string) {
	if len(qty) != len(hargaSatuan) {
		return false, "Salah"
	}

	for i := range qty {
		if qty[i] < 0 || hargaSatuan[i] < 0 {
			return false, "Salah"
		}
	}
	return true, ""
}

func TentukanStatus(total float64) string {
	switch {
	case total > 1000000:
		return "Prioritas"
	case total <= 1000000 && total > 100000:
		return "Reguler"
	default:
		return "Hemat"
	}
}

func RingkasanPesanan(qty []int, hargaSatuan []float64, tarifPajak float64) string {
	subtotal := HitungTotalPesanan(qty, hargaSatuan)
	diskon := HitungDiskon(subtotal)
	totalAkhir := TotalSetelahDiskon(qty, hargaSatuan, tarifPajak)
	status := TentukanStatus(totalAkhir)

	return fmt.Sprintf("Subtotal awalnya : %.2f\nDapat diskon: %.2f\nJadi total akhirnya: %.2f\n%s", subtotal, diskon, totalAkhir, status)
}

// TODO(Level 9): signature ini SUDAH benar (cari tahu sendiri kenapa
// bentuknya begini - lihat SOAL.md) - tinggal implementasikan isinya.
func Total(harga ...float64) float64 {
	panic("belum diimplementasikan")
}

// TODO(Level 10, bonus): signature ini SUDAH benar (cari tahu sendiri
// kenapa ada dua nilai balik - lihat SOAL.md) - tinggal implementasikan isinya.
func HitungOngkosKirim(beratKg float64, jarakKm float64) (float64, error) {
	panic("belum diimplementasikan")
}

func main() {
	fmt.Println("Sales Order Processor - pertemuan 2")
}
