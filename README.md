# フィボナッチ数を返すRESTAPIの開発

## 概要
指定されたn番目のフィボナッチ数をJSON形式でレスポンスとして返すRESTAPIである。(URL例：https://speee-api.onrender.com/fib?n=10)<br>
以下に詳細の仕様を示す。<br>

【仕様】<br>
エンドポイント：https://speee-api.onrender.com/fib<br>
クエリパラメータ：n={整数}<br>
メソッド：GETのみ<br>
レスポンスタイム：1秒まで<br>
エラー処理
* 入力値：0以下→エラーコード400
* 入力値：文字列→エラーコード400
* レスポンスタイム1秒以上→エラーコード504


<br>
【補足】<br>
無料サーバー使用のため、15分ほどAPIを利用していないとサーバーがスリープ状態になるため、初回API呼び出し時にはサーバーの起動時間により30秒ほどかかる可能性がある。<br>
そのため、APIの性能テストには初回を除いて確認していただきたい。