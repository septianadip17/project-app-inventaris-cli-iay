package entity

import "time"

type Category struct {
	ID          int
	Name        string
	Description string
}

type Item struct {
	ID           int
	CategoryID   int
	CategoryName string // Join result
	Name         string
	Price        float64
	PurchaseDate time.Time
	DaysUsed     int // Calculated field
}
