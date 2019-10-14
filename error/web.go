package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[len("/list/"):]

	if file, err := os.Open(path); err != nil {
		//panic(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	} else {
		defer file.Close()

		if all, err := ioutil.ReadAll(file); err != nil {
			panic(err)
		} else {
			writer.Write(all)
		}
	}
}

func main() {
	http.HandleFunc("/list/", handler)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(nil)
	}
}
