package account

import "github.com/cloudwego/hertz/pkg/app/server"

func customizedRegister(hertz *server.Hertz) {
	hertz.POST("/logup", LogUp)
	hertz.POST("/login", JwtMiddleware.LoginHandler)
	auth := hertz.Group("/auth", JwtMiddleware.MiddlewareFunc())
	auth.GET("/ping", Ping)
}

func Register(r *server.Hertz) {
	customizedRegister(r)
}
