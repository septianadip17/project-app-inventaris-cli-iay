package ui

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-iay/internal/entity"
	"text/tabwriter"
)

func PrintItems(items []entity.Item) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tNama Barang\tKategori\tHarga\tTgl Beli\tHari Pakai")
	fmt.Fprintln(w, "--\t-----------\t--------\t-----\t--------\t----------")

	for _, item := range items {
		fmt.Fprintf(w, "%d\t%s\t%s\tRp %.2f\t%s\t%d Hari\n",
			item.ID, item.Name, item.CategoryName, item.Price,
			item.PurchaseDate.Format("2006-01-02"), item.DaysUsed)
	}
	w.Flush()
}
