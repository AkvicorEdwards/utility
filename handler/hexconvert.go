package handler

import (
	"fmt"
	"net/http"
	"util/tpl"
)

func binaryCalculator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.BinaryCalculator.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}