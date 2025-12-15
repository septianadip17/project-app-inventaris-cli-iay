package main

import (
	"flag"
	"fmt"
	"project-app-inventaris-cli-iay/data"
	"project-app-inventaris-cli-iay/service"
	"project-app-inventaris-cli-iay/utils"
)

func main() {
	// 1. Init Database (Utils)
	db := utils.ConnectDB()
	defer db.Close()

	// 2. Init Layer (Dependency Injection)
	// Repository butuh DB, Service butuh Repository
	repo := data.NewRepository(db)
	svc := service.NewService(repo)

	// 3. Setup CLI Flags (Handler Input)
	mode := flag.String("mode", "", "Pilih mode: list-kategori, add-kategori")
	name := flag.String("name", "", "Nama untuk tambah data")
	desc := flag.String("desc", "", "Deskripsi untuk tambah data")

	flag.Parse()

	// 4. Routing Logic
	switch *mode {
	case "list-kategori":
		svc.ShowCategories() // Handler manggil Service
	case "add-kategori":
		svc.CreateCategory(*name, *desc) // Handler manggil Service
	default:
		fmt.Println("=== APLIKASI INVENTARIS ===")
		fmt.Println("Cara pakai:")
		fmt.Println("  go run main.go -mode=list-kategori")
		fmt.Println("  go run main.go -mode=add-kategori -name=\"Meja\" -desc=\"Furniture\"")
	}
}
