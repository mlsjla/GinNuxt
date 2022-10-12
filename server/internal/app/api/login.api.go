package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	jsoniter "github.com/json-iterator/go"

	"github.com/mlsjla/gin-nuxt/server/internal/app/config"
	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
	"github.com/mlsjla/gin-nuxt/server/internal/app/setting"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/logger"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/sms"
)

var LoginSet = wire.NewSet(wire.Struct(new(LoginAPI), "*"))

type LoginAPI struct {
	UserSrv     *service.UserSrv
	LoginSrv    *service.LoginSrv
	MenuSrv     *service.MenuSrv
	RoleMenuSrv *service.RoleMenuSrv
}

func (a *LoginAPI) WechatRegister(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		ginx.ResError(c, errors.New400Response("token not empty"))
		return
	}

	resp, err := http.Get("https://www.zhanmishu.com/plugin.php?id=zhanmishu_wechat:index&getUser=yes&token=" + token + "&openKey=a7da9591c7a0a4e82f407b086bbce32c")
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	userStr := jsoniter.Get(body, "data").ToString()

	ctx := c.Request.Context()
	var item schema.User

	item.ID = jsoniter.Get([]byte(userStr), "uid").ToUint64()
	if item.ID < 1 {
		ginx.ResError(c, errors.New400Response("no user"))
		return
	}
	dbUser, err := a.UserSrv.Get(ctx, item.ID)
	fmt.Println("dbUser", dbUser)
	if err == nil {
		tokenInfo, err := a.LoginSrv.GenerateToken(ctx, a.formatTokenUserID(dbUser.ID, dbUser.Username))
		if err != nil {
			ginx.ResError(c, err)
			return
		}
		

		ctx = logger.NewUserIDContext(ctx, dbUser.ID)
		ctx = logger.NewUserNameContext(ctx, dbUser.Username)
		ctx = logger.NewTagContext(ctx, "__login__")
		logger.WithContext(ctx).Infof("login")

		ginx.ResSuccess(c, tokenInfo)
		return
	}

	item.Username = jsoniter.Get([]byte(userStr), "username").ToString()
	item.Realname = item.Username
	item.Email = jsoniter.Get([]byte(userStr), "email").ToString()
	item.Status = 1
	item.Nickname = item.Username

	_, err = a.UserSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	tokenInfo, err := a.LoginSrv.GenerateToken(ctx, a.formatTokenUserID(item.ID, item.Username))
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, item.ID)
	ctx = logger.NewUserNameContext(ctx, item.Username)
	ctx = logger.NewTagContext(ctx, "__login__")
	logger.WithContext(ctx).Infof("login")

	ginx.ResSuccess(c, tokenInfo)
}

func (a *LoginAPI) SendSms(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.SmsRegisterParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	res := setting.Setting("sms")
	useCaptcha := jsoniter.Get([]byte(res), "useCaptcha").ToBool()

	if !schema.CheckIsRootUser(ctx, contextx.FromUserID(ctx)) && useCaptcha {
		if !captcha.VerifyString(item.CaptchaID, item.CaptchaCode) {
			ginx.ResError(c, errors.New400Response("无效的验证码"))
			return
		}
	}
	params := sms.SmsParams{
		Code: 11,
	}

	s, err := jsoniter.MarshalToString(params)
	if err != nil {
		ginx.ResError(c, err)
	}

	options := sms.SendParams{
		Mobile: item.Mobile,
		Value:  item.Value,
		Params: s,
	}
	send := sms.NewSMS().Send(options)
	ginx.ResSuccess(c, send)
}

func (a *LoginAPI) GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.LoginSrv.GetCaptcha(ctx, config.C.Captcha.Length)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *LoginAPI) ResCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	captchaID := c.Query("id")
	if captchaID == "" {
		ginx.ResError(c, errors.New400Response("captcha id not empty"))
		return
	}

	if c.Query("reload") != "" {
		if !captcha.Reload(captchaID) {
			ginx.ResError(c, errors.New400Response("not found captcha id"))
			return
		}
	}

	cfg := config.C.Captcha
	err := a.LoginSrv.ResCaptcha(ctx, c.Writer, captchaID, cfg.Width, cfg.Height)
	if err != nil {
		ginx.ResError(c, err)
	}
}

func (a *LoginAPI) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.LoginParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	if !captcha.VerifyString(item.CaptchaID, item.CaptchaCode) {
		ginx.ResError(c, errors.New400Response("无效的验证码"))
		return
	}
	user, err := a.LoginSrv.Verify(ctx, item.Username, item.Password)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	tokenInfo, err := a.LoginSrv.GenerateToken(ctx, a.formatTokenUserID(user.ID, user.Username))
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, user.ID)
	ctx = logger.NewUserNameContext(ctx, user.Username)
	ctx = logger.NewTagContext(ctx, "__login__")
	logger.WithContext(ctx).Infof("login")

	ginx.ResSuccess(c, tokenInfo)
}

func (a *LoginAPI) formatTokenUserID(userID uint64, userName string) string {
	return fmt.Sprintf("%d-%s", userID, userName)
}

func (a *LoginAPI) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	userID := contextx.FromUserID(ctx)
	if userID != 0 {
		ctx = logger.NewTagContext(ctx, "__logout__")
		err := a.LoginSrv.DestroyToken(ctx, ginx.GetToken(c))
		if err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
		logger.WithContext(ctx).Infof("logout")
	}
	ginx.ResOK(c)
}

func (a *LoginAPI) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	tokenInfo, err := a.LoginSrv.GenerateToken(ctx, a.formatTokenUserID(contextx.FromUserID(ctx), contextx.FromUserName(ctx)))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, tokenInfo)
}

func (a *LoginAPI) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	info, err := a.LoginSrv.GetLoginInfo(ctx, contextx.FromUserID(ctx))

	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, info)
}

func (a *LoginAPI) QueryUserMenuTree(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.MenuQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	// 根据roleMenu 获取ids
	// 获取role_id
	info, err := a.LoginSrv.GetLoginInfo(ctx, contextx.FromUserID(ctx))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	fmt.Println("info.Roles", info.Roles)
	result, err := a.MenuSrv.Query(ctx, params, schema.MenuQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC)),
	})
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResList(c, result.Data.ToTree())
}

func (a *LoginAPI) UpdateUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.UpdateInfoParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.LoginSrv.UpdateUserInfo(ctx, contextx.FromUserID(ctx), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
func (a *LoginAPI) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.UpdatePasswordParam
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.LoginSrv.UpdatePassword(ctx, contextx.FromUserID(ctx), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
