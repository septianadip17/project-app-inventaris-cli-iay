package main

import (
	"flag"
	"fmt"
	"os"

	// Import package internal kamu di sini
	// "project-app-inventaris-cli-iay/config"
	// "project-app-inventaris-cli-iay/internal/repository"
	// "project-app-inventaris-cli-iay/internal/ui"

	_ "github.com/lib/pq"
)

func main() {
	// 1. Setup Database
	// connStr := "user=postgres dbname=inventaris sslmode=disable password=secret"
	// db, err := sql.Open("postgres", connStr)
	// ... error handling ...

	// 2. Parsing Command
	// Contoh: ./app -mode=list-items
	// Contoh: ./app -mode=add-item -name="Meja" -price=500000

	mode := flag.String("mode", "", "Mode operasi: list-items, add-item, report, replacement")
	name := flag.String("name", "", "Nama barang/kategori")
	// ... tambahkan flag lain ...

	flag.Parse()

	switch *mode {
	case "list-items":
		// Panggil repository.GetAll()
		// Panggil ui.PrintItems()
	case "report":
		// Loop semua item
		// Panggil service.CalculateDepreciation()
		// Tampilkan total nilai aset
	case "replacement":
		// Filter item dengan DaysUsed > 100
	default:
		fmt.Println("Gunakan -help untuk melihat cara penggunaan")
		os.Exit(1)
	}
}
