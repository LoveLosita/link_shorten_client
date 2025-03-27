package init_client

import (
	"github.com/cloudwego/kitex/client"
	"link_shorten_client_and_port/link_gen/kitex_gen/link/linkservice"
)

var NewLinkClient linkservice.Client

func InitLinkSvClient() error {
	var err error
	NewLinkClient, err = linkservice.NewClient("linkservice", client.WithHostPorts("0.0.0.0:8890"))
	if err != nil {
		return err
	}
	return nil
}
