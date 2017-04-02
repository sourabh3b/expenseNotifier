package main

import (
	"fmt"
	"net/http"
	"github.com/sourabhbhagat/expenseNotifier/model"
	"github.com/unrolled/render"
	"encoding/json"
)

//SaveExpense -  saves expense
func SaveExpense(w http.ResponseWriter, r *http.Request){
	render := render.New()
	inputExpense := model.Expense{}
	err := json.NewDecoder(r.Body).Decode(&inputExpense)
	if err != nil {
		fmt.Println("Cannot decode expense ", err.Error())
		render.JSON(w, http.StatusBadRequest, "Cannot find respose Error : "+ err.Error())
		return
	}

	err = model.SaveExpense(inputExpense)
	if err != nil {
		fmt.Println("Cannot Save expense ", err.Error())
		render.JSON(w, http.StatusBadGateway, "Saved Expense ")
		return
	}
		fmt.Println("Saved expense ")
		render.JSON(w, http.StatusOK, "Saved Expense :")
		return
}
func main() {
	fmt.Println("Started expense notifier....")
	http.HandleFunc("/saveExpense", SaveExpense)
	http.ListenAndServe(":9753", nil)
}
