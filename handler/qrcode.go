package handler

import (
	"fmt"
	"net/http"
	"util/tpl"
)

func qrCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.QrCode.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

