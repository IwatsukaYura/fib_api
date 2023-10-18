package handlers

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"

	"github.com/IwatsukaYura/speee_api/models"
)

// n番目のフィボナッチ数を返すよ
// ただし、nは1以上の整数とする(0以下の場合は0を返す)
// func Fibonacci(n int) int {
// 	if n <= 0 {
// 		return 0
// 	}
// 	if n <= 1 {
// 		return n
// 	}

// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }

// var memo = make(map[uint64] uint64)

// func fibonacci(n uint64) uint64 {
// 	if n <= 0 {
// 		return 0
// 	}
// 	if n <= 1 {
// 		return 1
// 	}

// 	if val, ok := memo[n]; ok {
// 		return val
// 	}

// 	memo[n] = fibonacci(n-1) + fibonacci(n-2)
// 	return memo[n]
// }

func Fibonacci(n int64) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	a, b := big.NewInt(0), big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		a, b = b, new(big.Int).Add(a, b)
	}
	return b
}

// エラーをjson形式で返す
func jsonError(w http.ResponseWriter, message string) {
	Error := models.Error{Status: http.StatusBadRequest, Message: message}
	jsonError, err := json.Marshal(Error)
	if err != nil {
		http.Error(w, " 内部サーバーエラーです！", http.StatusInternalServerError)
		return
	}
	w.Write(jsonError)
}

func FibonacciHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var n int64
	var result *big.Int
	if p, ok := queryMap["n"]; ok && len(p) > 0 {
		var err error
		n, err = strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			jsonError(w, "クエリパラメータが正常ではありません。")
			return
		}
	} else {
		n = 1
	}

	result = Fibonacci(n)
	if result.Cmp(big.NewInt(0)) == 0 {
		jsonError(w, "1以上の整数を入力してください。")
		return
	}

	resultstruct := models.Result{Result: result}
	jsonFibonacciResult, err := json.Marshal(resultstruct)
	if err != nil {
		http.Error(w, " 内部サーバーエラーです！", http.StatusInternalServerError)
		return
	}

	w.Write(jsonFibonacciResult)
}
