package handler

import (
	"net/http"
	"fmt"
)

//GetExpense - Gets Expense as notification
//input : expenseID
//output : did you had breakfast/dinner ?
func GetExpense(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w , "hello")
}


//SaveExpense -  saves expense
func SaveExpense(w http.ResponseWriter, r *http.Request){

}