package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go createProxy("https://www.yahoo.com", "8085", "yahoo")
	wg.Add(1)
	go createProxy("https://www.google.com", "8086", "google")
	wg.Wait()
}

func createProxy(inputUrl string, port string, prefix string) {
	log.Println("Starting proxy for", inputUrl, prefix)
	remote, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.URL.Path = strings.Join(strings.Split(r.URL.String(), "/"+prefix)[1:], "")

			// remove prefix from r.URL
			// r.URL.Path = r.URL.Path[len("/"+prefix):]
			r.Host = remote.Host
			w.Header().Set("X-Ben", "Rad")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/"+prefix, handler(proxy))
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
