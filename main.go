package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloFunc(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello!")
}

func main() {
	http.HandleFunc("/", helloFunc)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}

}
