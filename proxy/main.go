package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	r := chi.NewRouter()
	port := "8080"
	proxy := NewReverseProxy("hugo", "1313")
	r.Use(middleware.Logger, middleware.Recoverer)
	r.Use(proxy.ReverseProxy)
	fmt.Println("The server is running on port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

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
		if strings.Contains(r.URL.Path, "api") {
			answer := []byte(fmt.Sprintln("Hello from API"))
			if _, err := w.Write(answer); err != nil {
				log.Fatal(err)
			}
			next.ServeHTTP(w, r)
			return
		}
		target := &url.URL{
			Scheme: "http",
			Host:   rp.host + ":" + rp.port,
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.Director = func(r *http.Request) {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.Host = target.Host
		}
		proxy.ServeHTTP(w, r)
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
