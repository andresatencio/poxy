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
