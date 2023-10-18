package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/IwatsukaYura/speee_api/models"
)

// n番目のフィボナッチ数を返すよ
// ただし、nは1以上の整数とする(0以下の場合は0を返す)
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
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

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var n int
	if p, ok := queryMap["n"]; ok && len(p) > 0 {
		var err error
		n, err = strconv.Atoi(p[0])
		if err != nil {
			jsonError(w, "クエリパラメータが正常ではありません。")
			return
		}
	} else {
		n = 1
	}

	result := fibonacci(n)
	if result == 0 {
		jsonError(w, "1以上の整数を入力してください。")
		return
	}
	test := models.Result{Result: result}
	jsonData, err := json.Marshal(test)
	if err != nil {
		http.Error(w, " 内部サーバーエラーです！", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
