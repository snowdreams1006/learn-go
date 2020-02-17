package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func handler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return errors.New("path must be start with " + prefix)
	}

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
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recover panic : %v", r)

				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		if err := handler(writer, request); err != nil {
			fmt.Printf("Error handling request with %s", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

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

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(handler))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(nil)
	}
}
