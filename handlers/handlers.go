package handlers

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/IwatsukaYura/speee_api/models"
)

// フィボナッチ数列のn番目の値を返す
// 引数：n番目を表す整数
// 戻り値：n番目のフィボナッチ数の値
// 例外：nが0以下の場合は0を返す
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
// 引数：http.ResponseWriter,ステータスコード ,エラーメッセージ
func jsonError(w http.ResponseWriter, status int ,message string) {
	Error := models.Error{Status: status, Message: message}
	jsonError, err := json.Marshal(Error)
	if err != nil {
		http.Error(w, " 内部サーバーエラーです！", http.StatusInternalServerError)
		return
	}
	w.Write(jsonError)
}

func FibonacciHandler(w http.ResponseWriter, req *http.Request) {
	var n int64         // n番目を表す整数
	var result *big.Int // n番目のフィボナッチ数の値

	// コンテキストを作成
	ctx, cancel := context.WithTimeout(req.Context(), time.Second) // タイムアウトを1秒に設定
	defer cancel()

	queryMap := req.URL.Query()
	if p, ok := queryMap["n"]; ok && len(p) > 0 {
		var err error
		n, err = strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			jsonError(w, http.StatusBadRequest,  "クエリパラメータが正常ではありません。")
			return
		}
	} else {
		jsonError(w, http.StatusBadRequest, "クエリパラメータが正常ではありません。")
		return
	}

	// Fibonacci関数をゴルーチンで呼び出し、タイムアウトを設定
	ch := make(chan *big.Int)
	go func() {
		result = Fibonacci(n)
		ch <- result
	}()

	select {
	case <-ctx.Done():
		jsonError(w, http.StatusGatewayTimeout, "処理がタイムアウトしました。")
		return
	case result = <-ch:
		// 0以下の整数が入力された場合はエラーを返す
		if result.Cmp(big.NewInt(0)) == 0 {
			jsonError(w, http.StatusBadRequest, "1以上の整数を入力してください。")
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
}
