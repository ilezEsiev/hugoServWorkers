package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("proxy started")
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", "1313")
	go WorkerTest()
	r.Use(proxy.ReverseProxy)
	r.Get("/api/*", ApiHandler)
	http.ListenAndServe(":8080", r)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
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
		if strings.HasPrefix(r.URL.Path, "/api/") {
			next.ServeHTTP(w, r)
			return
		}

		hugoURL, err := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
		if err != nil {
			log.Fatal(err)
		}
		httputil.NewSingleHostReverseProxy(hugoURL).ServeHTTP(w, r)
	})
}

func WorkerTest() {
	t := time.NewTicker(5 * time.Second)
	filePath := "/home/ilez/hugoproxy/hugo/content/tasks/_index.md"
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	n := strings.Index(string(file), "Счетчик: 0")

	var b int = 0
	for {
		select {
		case <-t.C:
			old := "Счетчик: 0"
			fmt.Println(b)
			new := fmt.Sprintf("Счетчик: %d", b+1)
			result := strings.Replace(string(file), old, new, n)

			err = os.WriteFile(filePath, []byte(result), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
