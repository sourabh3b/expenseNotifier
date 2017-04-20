package main

import (
	"encoding/json"
	"fmt"
	"github.com/sourabhbhagat/expenseNotifier/model"
	"github.com/unrolled/render"
	"net/http"
)

//SaveExpense -  saves expense
func SaveExpense(w http.ResponseWriter, r *http.Request) {
	render := render.New()
	inputExpense := model.Expense{}
	err := json.NewDecoder(r.Body).Decode(&inputExpense)
	if err != nil {
		fmt.Println("Cannot decode expense ", err.Error())
		render.JSON(w, http.StatusBadRequest, "Cannot find respose Error : "+err.Error())
		return
	}

	_,err = model.SaveExpense(inputExpense)
	if err != nil {
		fmt.Println("Cannot Save expense ", err.Error())
		render.JSON(w, http.StatusBadGateway, "Saved Expense ")
		return
	}
	fmt.Println("Saved expense ")
	render.JSON(w, http.StatusOK, "Saved Expense :")
	return
}

//GetExpense - handler to get expense by ID
func GetExpense(w http.ResponseWriter, r *http.Request) {
	render := render.New()

	expenseID := r.URL.Query().Get("expenseID")

	expenseObject,err := model.GetExpense(expenseID)
	if err != nil {
		fmt.Println("Cannot Save expense ", err.Error())
		render.JSON(w, http.StatusBadGateway, "Saved Expense ")
		return
	}
	render.JSON(w, http.StatusOK, expenseObject)
	return
}
func main() {
	fmt.Println("Started expense notifier....")
	http.HandleFunc("/saveExpense", SaveExpense)
	http.HandleFunc("/getExpense", GetExpense)
	http.ListenAndServe(":9753", nil)
}
