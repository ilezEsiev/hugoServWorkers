package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverseProxy(t *testing.T) {
	// Создаем новый Router и экземпляр ReverseProxy
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)
	r.Get("/api/*", ApiHandler)

	// Создаем тестовый сервер
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Тестируем запрос к /api/
	respApi, err := http.Get(ts.URL + "/api/")
	if err != nil {
		t.Fatal(err)
	}
	defer respApi.Body.Close()

	if respApi.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, respApi.StatusCode)
	}

	// Тестируем запрос к другим путям(404)
	respOtherBad, err := http.Get(ts.URL + "/some/path")
	if err != nil {
		t.Fatal(err)
	}
	defer respOtherBad.Body.Close()

	if respOtherBad.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, respOtherBad.StatusCode)
	}

	// Тестируем запрос к другим путям(404)
	respOtherGood, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer respOtherGood.Body.Close()

	if respOtherGood.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, respOtherGood.StatusCode)
	}
}

func TestNewReverseProxy(t *testing.T) {
	host := "example.com"
	port := "8080"

	proxy := NewReverseProxy(host, port)

	if proxy.host != host || proxy.port != port {
		t.Errorf("Expected host: %s, got host: %s; Expected port: %s, got port: %s", host, proxy.host, port, proxy.port)
	}
}

func TestApiHandler(t *testing.T) {
	// Создаем тестовый ResponseWriter и Request
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/", nil)

	// Вызываем ApiHandler с созданными ResponseWriter и Request
	ApiHandler(w, r)

	// Проверяем статус код и тело ответа
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expectedResponse := "Hello from API"
	if body := w.Body.String(); body != expectedResponse {
		t.Errorf("Expected response body: %s, got: %s", expectedResponse, body)
	}
}
