package main

import (
	"fmt"
	"net/http"
	"expenseNotifier/handler"
)

func main() {
	fmt.Println("Started expense notifier....")
	http.HandleFunc("/getExpense", handler.GetExpense)
	http.ListenAndServe(":8080", nil)
}
