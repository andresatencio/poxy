package poxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Poxy struct {
	mapUrlProx map[string]httputil.ReverseProxy
	middles    []http.HandlerFunc
}

func NewPoxy() *Poxy {
	var p Poxy
	p.mapUrlProx = map[string]httputil.ReverseProxy{}
	p.middles = []http.HandlerFunc{}
	return &p
}

func (p *Poxy) Add(in string, out string) {
	urlOut, err := url.Parse(out)
	if err != nil {
		panic(err)
	}

	prox := httputil.NewSingleHostReverseProxy(urlOut)
	p.mapUrlProx[in] = *prox
}

func (p *Poxy) Run() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		s := p.mapUrlProx[r.Host]
		for _, fn := range p.middles {
			fn(w, r)
		}
		s.ServeHTTP(w, r)
	}
}

func (p *Poxy) Middle(fn http.HandlerFunc) {
	p.middles = append(p.middles, fn)
}

// Eg. middleware
func mi2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Desde middle 2:" + r.Host)
}
