package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Sum(a int, b int) int {
	return a + b
}





func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello darling%d\n", Sum(1, 2))
}




















type response struct {
	ReqTime time.Time
	Value   string
	Error   string
}



func ApiCapitalize(w http.ResponseWriter, req *http.Request) {
	log.Println("got request for capitalization")
	r, err := ioutil.ReadAll(req.Body)
	var resp response
	if err != nil {
		w.WriteHeader(400)
		resp = response{
			ReqTime: time.Now(),
			Value:   "",
			Error:   err.Error(),
		}
		return
	}

	var val Request

	err = json.Unmarshal(r, &val)
	log.Println("given value: "+val.Value)
	if err != nil {
		w.WriteHeader(400)
		resp = response{
			ReqTime: time.Now(),
			Value:   "",
			Error:   err.Error(),
		}
		return
	}
	resp = response{
		ReqTime: time.Now(),
		Value:   strings.ToUpper(val.Value),
		Error:   "",
	}

	respJson, _ := json.Marshal(resp)
	w.Header().Set("Content-Type","application/json")

	log.Println("sending response")
	w.Write(respJson)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

var port string

func main() {


	ObtainPort()
	log.Println("starting")

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", home)
	http.HandleFunc("/api/capitalise", ApiCapitalize)

	fmt.Println(fmt.Sprintf("server started on port %s", port))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}

func ObtainPort() {
	defer recoverPort()

	port = os.Args[1]
}

func recoverPort() {
	if r := recover(); r != nil {
		fmt.Println("no port was given. set to default Value - 80 ", r)
		port = "80"
	}
}
