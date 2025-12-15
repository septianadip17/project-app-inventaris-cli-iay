package service

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-iay/data"
	"project-app-inventaris-cli-iay/model"
	"text/tabwriter"
)

type Service struct {
	Repo *data.Repository
}

func NewService(repo *data.Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateCategory(name, desc string) {
	// Validasi sederhana (Business Logic)
	if name == "" {
		fmt.Println("Error: Nama kategori tidak boleh kosong!")
		return
	}

	cat := model.Category{Name: name, Description: desc}
	err := s.Repo.AddCategory(cat)
	if err != nil {
		fmt.Println("Gagal menyimpan kategori:", err)
	} else {
		fmt.Println("Sukses menambahkan kategori:", name)
	}
}

func (s *Service) ShowCategories() {
	cats, err := s.Repo.GetAllCategories()
	if err != nil {
		fmt.Println("Error mengambil data:", err)
		return
	}

	// Tampilan Tabel (Sesuai soal)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tNama Kategori\tDeskripsi")
	fmt.Fprintln(w, "--\t-------------\t---------")

	for _, c := range cats {
		fmt.Fprintf(w, "%d\t%s\t%s\n", c.ID, c.Name, c.Description)
	}
	w.Flush()
}
