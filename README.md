# Poxy

A tiny lib for reverse proxy. Principal purpose is learn GoLang.

![Poxy](poxipol.png "Poxy")


## Install

```sh
go get -u github.com/andresatencio/poxy
```

## Examples

Let's start registering a couple of endpoint URL and internal app:

```go
func main() {
	poxy := poxy.NewPoxy()
	poxy.Add("test.fruta:3000", "http://localhost:5000")
	
	http.HandleFunc("/", poxy.Run())
	er := http.ListenAndServe(":3000", nil)

	if er != nil {
		fmt.Println(er)
		panic(er)
	}

}
```

HTTPS:

```go
package main

import (
	"crypto/tls"
	"net/http"

	"github.com/andresatencio/poxy"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	poxy := poxy.NewPoxy()
	poxy.Add("frutadev.com", "http://localhost:5000")
	poxy.Add("test.frutadev.com", "http://localhost:5001")

	mux := http.NewServeMux()
	mux.HandleFunc("/", poxy.Run())

	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache("~/.cache-letsencrypt"),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	server.ListenAndServeTLS("", "")
}

```
