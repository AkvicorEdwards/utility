package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

// 获取全局session
func GetUser(r *http.Request) (*sessions.Session, error) {
	return Get(r, "user")
}

// 设置用户的Session
func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	// Session
	ses, _ := GetUser(r)
	ses.Values["login"] = time.Now().UnixNano()
	ses.Options.MaxAge = 60 * 60 * 24 * 7
	err := ses.Save(r, w)
	if err != nil {
		_, _ = fmt.Fprintln(w, "ERROR session SetUserInfo")
		return
	}
}

func LoginTime(r *http.Request) int64 {
	ses, err := Get(r, "user")
	if err != nil {
		return 0
	}
	t, ok := ses.Values["login"].(int64)
	if !ok {
		return 0
	}
	return t
}
