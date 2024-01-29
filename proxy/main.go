package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"hugoproxy/proxy/graph"
	"hugoproxy/proxy/tree"
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
	go WorkerCounter()
	go graph.GenerateGraph()
	go tree.TreeWorker()
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

func WorkerCounter() {
	t := time.NewTicker(5 * time.Second)
	filePath := "/app/static/tasks/_index.md"
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	old := "Текущее время: 2021-10-13 15:00:00\n\nСчетчик: 0"
	n := strings.Index(string(file), old)

	var b int
	for {
		currentTime := time.Now()
		formattedTime := currentTime.Format("2006-01-02 15:04:05")
		select {
		case <-t.C:
			new := fmt.Sprintf("Текущее время: %s\n\nСчетчик: %d", formattedTime, b+1)
			result := strings.Replace(string(file), old, new, n)
			err = os.WriteFile(filePath, []byte(result), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
