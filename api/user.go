package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"link_shorten_client_and_port/init_client"
	"link_shorten_client_and_port/response"
	"link_shorten_client_and_port/user_gen/kitex_gen/user"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	var loginUser user.UserLoginRequest
	err := c.BindJSON(&loginUser)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	resp, _ := init_client.NewUserClient.UserLogin(context.Background(), &loginUser)
	/*if err != nil {
		c.JSON(consts.StatusInternalServerError, response.InternalError(err))
		return
	}*/
	c.JSON(consts.StatusOK, resp)
}

func UserRegister(ctx context.Context, c *app.RequestContext) {
	var registerUser user.UserRegisterRequest
	var emptyResp *user.UserRegisterResponse
	err := c.BindJSON(&registerUser)
	if err != nil {
		c.JSON(consts.StatusBadRequest, response.WrongParamType)
		return
	}
	resp, _ := init_client.NewUserClient.UserRegister(context.Background(), &registerUser)
	if resp != emptyResp {
		if resp.Status.Code == "500" {
			c.JSON(consts.StatusInternalServerError, resp.Status)
			return
		} else {
			c.JSON(consts.StatusBadRequest, resp.Status)
			return
		}
	}
	c.JSON(consts.StatusOK, response.Ok)
	return
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	//TODO

}
