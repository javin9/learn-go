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

type RequestUpdateBody struct {
	ID          int    `json:"id" binding:"required"` // 主键
	Title       string `json:"title"`                 // 标题
	Description string `json:"description"`           // 描述
	Author      string `json:"author"`                // 作者
	Thumbnail   string `json:"thumbnail"`             // 封面
	Category    int    `json:"category"`              // 分类
	Duration    int    `json:"duration"`              // 时长
	Resolution  string `json:"resolution"`            // 分辨率 例如： 720p 1080p
	FileSize    int64  `json:"file_size"`             // 文件大小 单位：字节
	Format      string `json:"format"`                // 文件格式 MP4、AVI
	Quality     int    `json:"quality"`               // 视频质量 1 高清，2标清
}

type ResponseUpdateBody struct {
	Message string `json:"message"`
}

func (me *CmsApp) ContentUpdate(ctx *gin.Context) {
	// 设置返回状态
	appResponse := app.AppResponse{GinContext: ctx}
	// 校验参数
	var req RequestUpdateBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if message, ok := err.(validator.ValidationErrors); ok {
			msg := validator_format.FormMessage(message)
			appResponse.ResponseError(e.INVALID_PARAMS, msg)
			return
		}

		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}

	server := dao.NewContentDao(me.mysqlDB)

	// 判断是否存在
	isExist, err := server.ExistByID(req.ID)
	if err != nil {
		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}

	if !isExist {
		appResponse.ResponseError(e.INVALID_PARAMS, "数据不存在")
		return
	}

	// // 更新
	if err := server.Update(req.ID, &model.Content{
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
	}); err != nil {
		appResponse.ResponseError(e.ERROR, err.Error())
		return
	}

	appResponse.ResponseSuccess("", "更新成功")
}
