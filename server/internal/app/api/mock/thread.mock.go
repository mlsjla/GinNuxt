package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ThreadSet = wire.NewSet(wire.Struct(new(ThreadMock), "*"))

type ThreadMock struct{}

// @Tags 文章
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {object} schema.ListResult{list=[]schema.Thread} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/threads [get]
func (a *ThreadMock) Query(c *gin.Context) {
}

// @Tags 文章
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.Thread
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/threads/{id} [get]
func (a *ThreadMock) Get(c *gin.Context) {
}

// @Tags 文章
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.Thread true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/threads [post]
func (a *ThreadMock) Create(c *gin.Context) {
}

// @Tags 文章
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.Thread true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/threads/{id} [put]
func (a *ThreadMock) Update(c *gin.Context) {
}

// @Tags 文章
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/threads/{id} [delete]
func (a *ThreadMock) Delete(c *gin.Context) {
}
