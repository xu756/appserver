package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func UploadRouter(r *route.RouterGroup) {
	r.POST("/file", UploadFile)
	r.POST("/file/list", UploadFileList)
}

// UploadFile 上传单个文件
func UploadFile(ctx context.Context, c *app.RequestContext) {

}

// UploadFileList 上传多个文件
func UploadFileList(ctx context.Context, c *app.RequestContext) {

}
