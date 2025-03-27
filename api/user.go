package api

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"link_shorten_client_and_port/init_client"
	"link_shorten_client_and_port/response"
	"link_shorten_client_and_port/user_gen/kitex_gen/user"
	"time"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var loginUser user.UserLoginRequest
	//1.从请求中获取参数
	err := c.BindJSON(&loginUser)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewUserClient.UserLogin(ctx, &loginUser)
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
	//3.返回结果
	c.JSON(consts.StatusOK, resp)
}

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var registerUser user.UserRegisterRequest
	//1.从请求中获取参数
	err := c.BindJSON(&registerUser)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewUserClient.UserRegister(ctx, &registerUser)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(consts.StatusInternalServerError, response.RpcServerConnectTimeOut)
		} else {
			c.JSON(consts.StatusInternalServerError, response.InternalError(err))
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

	//3.返回结果
	c.JSON(consts.StatusOK, response.Ok)
	return
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	var refreshTokenRequest user.TokenRefreshRequest
	//1.获取refreshToken
	err := c.BindJSON(&refreshTokenRequest)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	//2.调用服务端接口
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //设置超时时间
	defer cancel()
	resp, err := init_client.NewUserClient.TokenRefresh(ctx, &refreshTokenRequest)
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
	//3.返回结果
	c.JSON(consts.StatusOK, resp)
}
