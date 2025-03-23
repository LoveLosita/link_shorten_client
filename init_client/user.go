package init_client

import (
	"github.com/cloudwego/kitex/client"
	"link_shorten_client_and_port/user_gen/kitex_gen/user/userservice"
)

var NewUserClient userservice.Client

func InitUserClient() error {
	var err error
	NewUserClient, err = userservice.NewClient("user_login", client.WithHostPorts("0.0.0.0:8889"))
	if err != nil {
		return err
	}
	return nil
}
