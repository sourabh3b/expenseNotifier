package model

import "time"

//Expense - data model
type Expense struct {
	ID          string    `bson:"_id" json:"id"`
	ExpenseTime time.Time `bson:"expenseTime" json:"expenseTime"`
	IsLunch     bool      `bson:"isLunch" json:"isLunch"`
	Total       float32   `bson:"total" json:"total"`
}

