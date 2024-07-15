package controller

import (
	"github.com/gogf/gf/net/ghttp"
	"go-upload-img/app/model"
	"go-upload-img/app/tools"
)

// 上传图片 以base64的方式上传图片
func Upload_Base64(r *ghttp.Request) {
	var mUpload *model.Upload_Base64
	if err := r.Parse(&mUpload); err != nil {
		tools.Return_Status(r, nil, "400", err.Error())
		return
	}
	tools.Base64_To_Pic(r, mUpload.Img, "", "")
	return
}

// 上传文件
func Upload_File(r *ghttp.Request) {
	var mUpload *model.Upload_File //struct的file无法使用，后期补充
	if err := r.Parse(&mUpload); err != nil {
		tools.Return_Status(r, nil, "400", err.Error())
		return
	}
	tools.File_To_Pic(r, r.GetUploadFile("file"), mUpload.FileName, "")
	return
}
