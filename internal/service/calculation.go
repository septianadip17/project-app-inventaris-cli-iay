package service

import (
	"math"
	"time"
)

// CalculateDepreciation menghitung nilai sekarang berdasarkan saldo menurun
func CalculateDepreciation(price float64, purchaseDate time.Time) (float64, float64) {
	rate := 0.20 // 20%

	// Hitung selisih tahun (bisa pecahan)
	duration := time.Since(purchaseDate).Hours() / 24 / 365

	// Rumus Saldo Menurun: Value = Price * (1 - rate)^n
	currentValue := price * math.Pow((1-rate), duration)

	depreciationAmount := price - currentValue

	return currentValue, depreciationAmount
}

// IsNeedsReplacement cek jika barang sudah dipakai > 100 hari
func IsNeedsReplacement(purchaseDate time.Time) bool {
	days := time.Since(purchaseDate).Hours() / 24
	return days > 100
}
