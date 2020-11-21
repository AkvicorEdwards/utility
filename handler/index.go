package handler

import (
	"fmt"
	"net/http"
	"util/def"
	"util/session"
	"util/tpl"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Index.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Login.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != def.Account || password != def.Password {
		_, _ = fmt.Fprintln(w, "用户不存在或密码错误")
		return
	}

	session.SetUserInfo(w, r)

	http.Redirect(w, r, "/", 302)

}