# Poxi

A tiny lib for reverse proxy. Principal purpose is learn GoLang.

![Poxy](poxipol.png "Poxy")


## Install

```sh
go get -u github.com/andresatencio/poxi
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
