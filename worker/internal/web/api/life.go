package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./view/home.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	f, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(f))
}
