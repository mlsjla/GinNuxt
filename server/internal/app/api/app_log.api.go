package api

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/mlsjla/gin-nuxt/server/pkg/util/json"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/contextx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/ginx"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/internal/app/service"
)

var AppLogSet = wire.NewSet(wire.Struct(new(AppLogAPI), "*"))
var shellPath = "/Users/mac/test"

type AppLogAPI struct {
	AppLogSrv *service.AppLogSrv
	AppSrv    *service.AppSrv
}

func (a *AppLogAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.AppLogQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.AppLogSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *AppLogAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.AppLogSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *AppLogAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.AppLog
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}
	UserId := contextx.FromUserID(ctx)
	if UserId == 0 {
		ginx.ResError(c, errors.New("用户数据异常"))
		return
	}
	item.UserId = UserId

	result, err := a.AppLogSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *AppLogAPI) Preview(c *gin.Context) {

	ctx := c.Request.Context()
	app, err := a.AppSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	UserId := contextx.FromUserID(ctx)
	if UserId != app.UserId {
		ginx.ResError(c, errors.New("用户数据异常"))
		return
	}
	// 获取证书配置信息等

	var item schema.AppPack
	json.Unmarshal([]byte(app.Data), &item)
	item.Appid = app.Appid

	// 复制秘钥文件
	// config.C.Local.Path
	_, err = copyFile(shellPath+"/keys/private."+item.Appid+".key", item.Privatekey)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	// 开始执行命令进行预览

	var cmd *exec.Cmd
	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("sh", shellPath+"/create.sh", item.Appid, item.Api, "preview")
	cmd.Wait()
	if _, err = cmd.CombinedOutput(); err != nil {
		os.Exit(1)
		return
	}
	ff, _ := ioutil.ReadFile(shellPath + "/preview/" + item.Appid + ".jpg")
	bufstore := make([]byte, 5000000)       //数据缓存
	base64.StdEncoding.Encode(bufstore, ff) // 文件转base64

	ginx.ResSuccess(c, ff)
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func (a *AppLogAPI) Upload(c *gin.Context) {
	ctx := c.Request.Context()
	app, err := a.AppSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	UserId := contextx.FromUserID(ctx)
	if UserId != app.UserId {
		ginx.ResError(c, errors.New("用户数据异常"))
		return
	}
	// 获取证书配置信息等

	var item schema.AppPack
	json.Unmarshal([]byte(app.Data), &item)
	item.Appid = app.Appid

	// 复制秘钥文件
	// config.C.Local.Path
	_, err = copyFile(shellPath+"/keys/private."+item.Appid+".key", item.Privatekey)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	// 开始执行命令进行预览

	var upload []byte
	var cmd *exec.Cmd
	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("sh", shellPath+"/create.sh", item.Appid, item.Api, "upload")
	cmd.Wait()
	if upload, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ginx.ResSuccess(c, string(upload))
}

func (a *AppLogAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.AppLog
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.AppLogSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *AppLogAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.AppLogSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
