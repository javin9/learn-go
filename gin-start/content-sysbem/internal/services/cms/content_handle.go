package cms

import (
	"content-sysbem/internal/app"
	"content-sysbem/internal/dao"
	"content-sysbem/internal/model"
	"content-sysbem/internal/validator_format"
	"content-sysbem/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RequestBody struct {
	Title       string `json:"title" binding:"required"`      // 标题
	Description string `json:"description"`                   // 描述
	Author      string `json:"author" binding:"required"`     // 作者
	Thumbnail   string `json:"thumbnail" binding:"required"`  // 封面
	Category    int    `json:"category" binding:"required"`   // 分类
	Duration    int    `json:"duration"`                      // 时长
	Resolution  string `json:"resolution" binding:"required"` // 分辨率 例如： 720p 1080p
	FileSize    int64  `json:"file_size"`                     // 文件大小 单位：字节
	Format      string `json:"format" binding:"required"`     // 文件格式 MP4、AVI
	Quality     int    `json:"quality" binding:"required"`    // 视频质量 1 高清，2标清
}

type ResponseBody struct {
	ID int `json:"id"` // 主键
}

// ContentController 添加内容
func (me *CmsApp) ContentCreate(ctx *gin.Context) {
	// 设置返回状态
	appResponse := app.AppResponse{GinContext: ctx}

	// 校验参数
	var req RequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if message, ok := err.(validator.ValidationErrors); ok {
			msg := validator_format.FormMessage(message)
			appResponse.ResponseError(e.INVALID_PARAMS, msg)
			return
		}

		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}

	// server 层
	contentDao := dao.NewContentDao(me.mysqlDB)
	content := model.Content{
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		Thumbnail:   req.Thumbnail,
		Category:    req.Category,
		Duration:    req.Duration,
		Resolution:  req.Resolution,
		FileSize:    req.FileSize,
		Format:      req.Format,
		Quality:     req.Quality,
	}
	if err := contentDao.Create(&content); err != nil {
		appResponse.ResponseError(e.ERROR, "添加内容失败")
		return
	}

	appResponse.ResponseSuccess(ResponseBody{ID: content.ID}, "添加内容成功")
}
