package main

import (
	"fmt"
	"net/http"
	"github.com/sourabhbhagat/expenseNotifier/handler"
)

func main() {
	fmt.Println("Started expense notifier....")
	http.HandleFunc("/getExpense", handler.GetExpense)
	http.ListenAndServe(":8080", nil)
}
