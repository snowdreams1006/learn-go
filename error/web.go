package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]

	if file, err := os.Open(path); err != nil {
		//panic(err)
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	} else {
		defer file.Close()

		if all, err := ioutil.ReadAll(file); err != nil {
			//panic(err)
			return err
		} else {
			writer.Write(all)

			return nil
		}
	}
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := handler(writer, request); err != nil {
			fmt.Printf("Error handling request with %s", err.Error())

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(handler))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(nil)
	}
}
