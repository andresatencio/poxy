package main

import (
	"fmt"
	"net/http"

	"github.com/andresatencio/poxy"
)

func middleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Desde middle 2:" + r.Host)
}

func main() {
	poxy := poxy.NewPoxy()

	poxy.Add("test.fruta:3000", "http://localhost:5000")
	poxy.Add("fruta:3000", "http://localhost:6000")
	poxy.Middle(middleware)

	http.HandleFunc("/", poxy.Run())
	er := http.ListenAndServe(":3000", nil)

	if er != nil {
		fmt.Println(er)
		panic(er)
	}

}
