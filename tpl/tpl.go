package tpl

import (
	"fmt"
	"html/template"
	"log"
)

var err error
var errCnt = 0

var Login *template.Template
var Index *template.Template
var Timestamp *template.Template
var TcpPort *template.Template
var QrCode *template.Template
var Colour *template.Template
var HexCalculator *template.Template
var BinaryCalculator *template.Template
var AgeCalculator *template.Template
var HexConvert *template.Template
var Timezone *template.Template
var JsPopUpWindow *template.Template

func Parse() {
	err = nil
	errCnt = 0

	Login = addFromFile("./tpl/login.html")
	Index = addFromFile("./tpl/index.html")
	Timestamp = addFromFile("./tpl/timestamp.html")
	TcpPort = addFromFile("./tpl/tcpport.html")
	QrCode = addFromFile("./tpl/qrcode.html")
	Colour = addFromFile("./tpl/colour.html")
	HexCalculator = addFromFile("./tpl/hex-calculator.html")
	BinaryCalculator = addFromFile("./tpl/binary-calculator.html")
	AgeCalculator = addFromFile("./tpl/age-calculator.html")
	HexConvert = addFromFile("./tpl/hexconvert.html")
	Timezone = addFromFile("./tpl/timezone.html")
	JsPopUpWindow = addFromFile("./tpl/javascript-pop-up-window.html")

	log.Printf("Parsing the html template was completed with %d errors\n", errCnt)
}

func add(name, tpl string) (t *template.Template) {
	t, err = template.New(name).Parse(tpl)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}

func addFromFile(file string) (t *template.Template) {
	t, err = template.ParseFiles(file)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}
