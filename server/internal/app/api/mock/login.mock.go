package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var LoginSet = wire.NewSet(wire.Struct(new(LoginMock), "*"))

type LoginMock struct {
}

// @Tags LoginAPI
// @Summary 获取验证码信息
// @Success 200 {object} schema.LoginCaptcha
// @Router /api/v1/pub/login/captchaid [get]
func (a *LoginMock) GetCaptcha(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 响应图形验证码
// @Param id query string true "验证码ID"
// @Param reload query string false "重新加载"
// @Produce image/png
// @Success 200 "图形验证码"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/pub/login/captcha [get]
func (a *LoginMock) ResCaptcha(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 用户登录
// @Param body body schema.LoginParam true "请求参数"
// @Success 200 {object} schema.LoginTokenInfo
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/pub/login [post]
func (a *LoginMock) Login(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 用户登出
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Router /api/v1/pub/login/exit [post]
func (a *LoginMock) Logout(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 刷新令牌
// @Security ApiKeyAuth
// @Success 200 {object} schema.LoginTokenInfo
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/pub/refresh-token [post]
func (a *LoginMock) RefreshToken(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 获取当前用户信息
// @Security ApiKeyAuth
// @Success 200 {object} schema.UserLoginInfo
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/pub/current/user [get]
func (a *LoginMock) GetUserInfo(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 查询当前用户菜单树
// @Security ApiKeyAuth
// @Success 200 {object} schema.ListResult{list=[]schema.MenuTree} "查询结果"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/pub/current/menutree [get]
func (a *LoginMock) QueryUserMenuTree(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 更新个人密码
// @Security ApiKeyAuth
// @Param body body schema.UpdatePasswordParam true "请求参数"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/pub/current/password [put]
func (a *LoginMock) UpdatePassword(c *gin.Context) {
}

// @Tags LoginAPI
// @Summary 更新个人信息
// @Security ApiKeyAuth
// @Param body body schema.UpdateInfoParam true "请求参数"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/pub/current/user [put]
func (a *LoginMock) UpdateUserInfo(c *gin.Context) {
}
