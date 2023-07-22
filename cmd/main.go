package main

import (
	"demo-hertz/internal/app/account"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {

	account.Init()
	account.InitJwt()

	h := server.Default()

	account.Register(h)
	h.Spin()
}
