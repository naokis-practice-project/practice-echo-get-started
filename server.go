package main

// https://github.com/labstack/echo
import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", hello)

	// Fatalを呼び出すとスタックトレースを出力してプロセスは終了
	e.Logger.Fatal(e.Start(":1323"))

}

// echo.Context 現在の HTTP リクエストのコンテキスト
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
