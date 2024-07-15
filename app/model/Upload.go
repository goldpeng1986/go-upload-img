package model

import "github.com/gogf/gf/net/ghttp"

type Upload_Base64 struct {
	Img     string `json:"img" v:"required#图片不能为空"` //base64图片
	ImgName string `json:"img_name,omitempty"`      //图片名称
	ImgPath string `json:"img_path,omitempty"`      //图片路径
}
type Upload_File struct {
	File     *ghttp.UploadFile `json:"file"`                //文件
	FileName string            `json:"file_name,omitempty"` //文件名称
	FilePath string            `json:"file_path,omitempty"` //文件路径
}
