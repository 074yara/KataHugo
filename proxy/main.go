package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	r := chi.NewRouter()
	proxy := NewReverseProxy(":", "1313")
	r.Use(proxy.ReverseProxy)

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := &url.URL{
			Scheme: "http",
			Host:   rp.host + ":" + rp.port,
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.Director = func(r *http.Request) {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
		}
		next.ServeHTTP(w, r)
	})
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
