package main

import (
	"flag"
	"log"
	"net/http"
	"module"
)

var addr = flag.String("addr","127.0.0.1:1718","Rest Service")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc  {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "{\"code\":500, \"desc\":\"" + r.Method + " not allowed.\"}", http.StatusBadRequest)
		} else {
			log.Println("Before go to fb")
			fn(w,r)
		}
	}
}

func main(){
	flag.Parse()
	log.Println("Start Server ..")

	http.HandleFunc("/rest/parser",makeHandler(module.JSONParserHandler))
	http.ListenAndServe(*addr,nil)
}
