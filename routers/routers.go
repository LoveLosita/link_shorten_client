package routers

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"link_shorten_client_and_port/api"
)

func RegisterRouters() {
	h := server.Default()

	userGroup := h.Group("/user")
	linkGroup := h.Group("/link")
	h.GET("/:id", api.LinkRedirect)

	userGroup.GET("/login", api.UserLogin)
	userGroup.POST("/register", api.UserRegister)
	userGroup.GET("/refresh_token", api.RefreshToken)

	linkGroup.POST("/generate", api.GenerateLink)

	h.Spin()
}
