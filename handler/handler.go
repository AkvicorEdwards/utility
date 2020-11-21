package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"util/session"
)

type str2func map[string]func(http.ResponseWriter, *http.Request)

var public str2func
var private str2func

func ParsePrefix() {
	public = make(str2func)
	private = make(str2func)

	private["/"] = index

	public["/login"] = login
	public["/myip"] = myIp
	public["/timestamp"] = unixTimestamp
	public["/tcpport"] = tcpPort
	public["/qrcode"] = qrCode
	public["/colour"] = colour
	public["/hexcalculator"] = hexCalculator
	public["/binarycalculator"] = binaryCalculator
	public["/agecalculator"] = ageCalculator
	public["/hexconvert"] = hexConvert
	public["/timezone"] = timezone
	public["/jspopupwindow"] = jsPopUpWindow
}

type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if resource(w, r, "/img/", "favicon.ico") {
		return
	}
	if resource(w, r, "/img/", "back.png") {
		return
	}
	if resource(w, r, "/js/", "jquery.min.js") {
		return
	}
	if resource(w, r, "/js/", "jquery-qrcode-0.14.0.js") {
		return
	}
	if resource(w, r, "/js/", "scripts.js") {
		return
	}
	if resource(w, r, "/css/", "all-7bb820b4.css") {
		return
	}
	if resource(w, r, "/css/", "main.css") {
		return
	}
	if resource(w, r, "/js/", "analytics.js") {
		return
	}
	if resource(w, r, "/js/", "body-0c9526df.js") {
		return
	}
	if resource(w, r, "/js/", "linkid.js") {
		return
	}
	if resource(w, r, "/js/", "lxe4nsf.js") {
		return
	}
	if resource(w, r, "/js/", "mixpanel-2-latest.min.js") {
		return
	}
	if resource(w, r, "/css/", "base.css") {
		return
	}
	if resource(w, r, "/css/", "mpublic_v2.css") {
		return
	}
	if resource(w, r, "/css/", "number.css") {
		return
	}
	if resource(w, r, "/js/", "common.js") {
		return
	}
	if resource(w, r, "/js/", "jquery-1.7.2.js") {
		return
	}
	if resource(w, r, "/css/", "basic.css") {
		return
	}
	if resource(w, r, "/css/", "bootstrap.min.css") {
		return
	}
	if resource(w, r, "/js/", "javascript-pop-up-window2.js") {
		return
	}

	if h, ok := public[r.URL.Path]; ok {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		h(w, r)
		return
	}

	if session.LoginTime(r) <= 0 {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		public["/login"](w, r)
		return
	}

	if h, ok := private[r.URL.Path]; ok {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		h(w, r)
		return
	}
}

func resource(w http.ResponseWriter, r *http.Request, path string, file string) bool {
	if path == "/css/" {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	} else if path == "/js/" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	} else if path == "/img/" {
		w.Header().Set("Content-Type", "image/png; charset=utf-8")
	}
	if ok, _ := regexp.MatchString(file, r.URL.Path); ok {
		download(w, "./tpl" + path + file)
		return true
	}
	return false
}

func download(w http.ResponseWriter, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		_, _ = fmt.Fprintln(w, "File Not Found")
		return
	}
	defer func() { _ = file.Close() }()
	data := make([]byte, 1024)
	for {
		n, err1 := file.Read(data)
		if err1 != nil && err1 != io.EOF {
			_, _ = fmt.Fprintln(w, "File Read Error")
			return
		}
		nn, err2 := w.Write(data[:n])
		if err2 != nil || nn != n {
			_, _ = fmt.Fprintln(w, "File Write Error")
			return
		}
		if err1 == io.EOF {
			return
		}
	}
}

func Fprint(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}

func Fprintf(w http.ResponseWriter, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}

func Fprintln(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprintln(w, a...)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
