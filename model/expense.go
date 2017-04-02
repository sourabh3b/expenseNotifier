package model

import(
	"errors"
	"fmt"
	"time"
	"gopkg.in/mgo.v2"
)

//Expense - data model
type Expense struct {
	ExpenseTime time.Time `bson:"expenseTime" json:"expenseTime"`
	IsLunch     bool      `bson:"isLunch" json:"isLunch"`
	Total       float32   `bson:"total" json:"total"`
}


//SaveExpense - handler to save expenses
func SaveExpense(expenseObject Expense) (error){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Mongo error", err.Error())
		return errors.New("Mongo connection Error "+ err.Error())
	}

	defer session.Close()

	// Collection Expense
	c := session.DB("test").C("Expense")

	expenseObject.ExpenseTime = time.Now()
	err = c.Insert(expenseObject)
	if err != nil {
		fmt.Println("DB insert error", err.Error())
		return errors.New("Cannot insert data into DB "+ err.Error())
	}
	return err
}