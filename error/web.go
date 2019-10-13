package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]

		if file, err := os.Open(path); err != nil {
			panic(err)
		} else {
			defer file.Close()

			if all, err := ioutil.ReadAll(file); err != nil {
				panic(err)
			} else {
				writer.Write(all)
			}
		}
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(nil)
	}
}
