package api

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"link_shorten_client_and_port/init_client"
	"link_shorten_client_and_port/link_gen/kitex_gen/link"
	"link_shorten_client_and_port/response"
	"time"
)

func GenerateLink(ctx context.Context, c *app.RequestContext) {
	var generateLinkRequest link.GenerateLinkRequest
	// 1.获取用户的token（如果有）
	accessToken := c.GetHeader("Authorization")
	strToken := string(accessToken)
	if string(accessToken) != "" {
		generateLinkRequest.Token = &strToken
	}
	// 2.获取用户的url
	err := c.BindJSON(&generateLinkRequest)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	// 2.1.检查参数的合法性
	if generateLinkRequest.LongUrl == "" {
		c.JSON(consts.StatusBadRequest, response.MissingParam)
		return
	}
	// 3.调用短链接服务的生成短链接接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewLinkClient.GenerateLink(ctx, &generateLinkRequest)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(consts.StatusInternalServerError, response.RpcServerConnectTimeOut)
			return
		} else {
			c.JSON(consts.StatusInternalServerError, response.InternalError(err))
			return
		}
	}
	if resp != nil {
		if resp.Status.Code == "500" { //如果是内部错误
			c.JSON(consts.StatusInternalServerError, resp.Status)
			return
		} else if resp.Status.Code != "10000" { //如果是参数错误
			c.JSON(consts.StatusBadRequest, resp.Status)
			return
		}
	}
	// 4.返回短链接
	c.JSON(consts.StatusOK, resp)
}

func LinkRedirect(ctx context.Context, c *app.RequestContext) {
	//1.获取用户路由中的短链
	shortLink := c.Param("id")
	//2.调用服务端接口，获取实际链接
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewLinkClient.LinkRedirect(ctx, &link.LinkRedirectRequest{ShortUrl: shortLink})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(consts.StatusInternalServerError, response.RpcServerConnectTimeOut)
			return
		} else {
			c.JSON(consts.StatusInternalServerError, response.InternalError(err))
			return
		}
	}
	if resp != nil {
		if resp.Status.Code == "500" {
			c.JSON(consts.StatusInternalServerError, resp.Status)
			return
		} else if resp.Status.Code != "10000" { //如果是参数错误
			c.JSON(consts.StatusBadRequest, resp.Status)
			return
		}
	}
	//3.若正常，将用户重定向至目标链接；若不正常，则返回错误
	if resp != nil {
		if *resp.LongUrl == "" {
			c.JSON(consts.StatusBadRequest, resp.Status)
			return
		}
		c.Redirect(consts.StatusFound, []byte(*resp.LongUrl))
		return
	}
}
