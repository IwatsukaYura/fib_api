package main

import (
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IwatsukaYura/fib_api/handlers"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int64
		expected *big.Int
	}{
		{0, big.NewInt(0)},
		{1, big.NewInt(1)},
		{5, big.NewInt(5)},
		{10, big.NewInt(55)},
		{20, big.NewInt(6765)},
	}

	for _, test := range tests {
		result := handlers.Fibonacci(test.n)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("Fibonacci(%d) = %v, want %v", test.n, result, test.expected)
		}
	}
}

func TestFibonacciHandler(t *testing.T) {
	tests := []struct {
		query    string
		expected string
	}{
		{"?n=5", `{"result":5}`},
		{"?n=10", `{"result":55}`},
		{"?n=0", `{"status":400,"message":"1以上の整数を入力してください。"}`},
		{"", `{"status":400,"message":"クエリパラメータが正常ではありません。"}`},
		{"?n=abc", `{"status":400,"message":"クエリパラメータが正常ではありません。"}`},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/fibonacci"+test.query, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.FibonacciHandler)

		handler.ServeHTTP(rr, req)

		if rr.Body.String() != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), test.expected)
		}
	}
}
