package handler

import (
	"net/http"
	"fmt"
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
		fmt.Println("Cannot find expense ", err.Error())
		render.JSON(w, http.StatusOK, "Saved Expense " + inputExpense.ID)
		return
	}
}