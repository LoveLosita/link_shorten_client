package cmd

import (
	"link_shorten_client_and_port/init_client"
	"link_shorten_client_and_port/routers"
	"log"
)

func Start() {
	//1.启动kitex客户端
	err := init_client.InitUserLoginClient()
	if err != nil {
		log.Fatal(err)
	}
	/*err = init_client.InitUserRegisterClient()
	if err != nil {
		log.Fatal(err)
	}*/
	//2.启动路由
	routers.RegisterRouters()
}
