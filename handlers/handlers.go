package handlers

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"

	"github.com/IwatsukaYura/speee_api/models"
)


// フィボナッチ数列のn番目の値を返す
//引数：n番目を表す整数
//戻り値：n番目のフィボナッチ数の値
//例外：nが0以下の場合は0を返す
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

// エラーをjson形式で返す関数
//引数：http.ResponseWriter, エラーメッセージ
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
	var n int64		// n番目を表す整数
	var result *big.Int	// n番目のフィボナッチ数の値
	
	queryMap := req.URL.Query()
	if p, ok := queryMap["n"]; ok && len(p) > 0 {
		var err error
		n, err = strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			jsonError(w, "クエリパラメータが正常ではありません。")
			return
		}
	} else {
		jsonError(w, "クエリパラメータが正常ではありません。")
		return
	}

	result = Fibonacci(n)

	// 0以下の整数が入力された場合はエラーを返す
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
