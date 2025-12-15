package service

import (
	"testing"
	"time"
)

func TestCalculateDepreciation(t *testing.T) {
	price := 1000000.0
	// Anggap beli 1 tahun lalu
	buyDate := time.Now().AddDate(-1, 0, 0)

	currentVal, _ := CalculateDepreciation(price, buyDate)

	// Harapannya sisa 80% (karena depresiasi 20%)
	expected := 800000.0

	// Toleransi komparasi float
	if currentVal < expected-1000 || currentVal > expected+1000 {
		t.Errorf("Expected about %.2f, got %.2f", expected, currentVal)
	}
}
