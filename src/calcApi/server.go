package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

var (
	operators = map[string]rune{
		"plus":     '+',
		"minus":    '-',
		"multiply": '*',
		"divide":   '/',
		"exp":      '^',
		"rad":      '|',
	}
)

func ftoa(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

func calc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<!DOCTYPE html><html><head>")
	fmt.Fprintf(w, "<style> body { height: 200px; background-color: red; background-image: linear-gradient(to right, red , yellow); } </style>")
	fmt.Fprintf(w, "<style> h1{ font-family: 'Inconsolata', monospace; color: #FFFFFF; font-size: 35pt;	}</style>")
	fmt.Fprintf(w, "</head>")
	fmt.Fprintf(w, "<body>")

	query := r.URL.RawQuery

	u, err := url.ParseQuery(query)
	if err != nil {
		log.Fatal(err)
	}

	a, erra := strconv.ParseFloat(u["a"][0], 64)
	b, errb := strconv.ParseFloat(u["b"][0], 64)
	op := u["operator"][0]

	resp := fmt.Sprintf("%f %c %f = ", a, operators[op], b)

	switch {
	case op == "plus":
		resp += ftoa(a + b)
	case op == "minus":
		resp += ftoa(a - b)
	case op == "multiply":
		resp += ftoa(a * b)
	case op == "divide":
		resp += ftoa(a / b)
	case op == "exp":
		resp += ftoa(math.Pow(a, b))
	case op == "rad":
		resp += ftoa(math.Pow(a, 1./b))

	default:
		resp = "unknown operator"
	}

	if erra != nil {
		resp = "Invalid format of a!"
	} else if errb != nil {
		resp = "Invalid format of b!"
	}

	fmt.Fprintf(w, "<h1>"+resp+"</h1>")
	fmt.Fprintf(w, "</body></html>")
}

func help(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("calcApi/help.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(file))
}

func main() {
	http.HandleFunc("/calc", calc)
	http.HandleFunc("/help", help)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
