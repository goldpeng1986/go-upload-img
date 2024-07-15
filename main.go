package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-upload-img/app/controller"
)

func main() {
	//开启服务器
	s := g.Server()
	//静态目录
	s.SetIndexFolder(true)
	s.AddSearchPath(g.Cfg().GetString("Domain.img_dir"))
	s.AddSearchPath(g.Cfg().GetString("Domain.file_dir"))
	s.AddStaticPath(g.Cfg().GetString("Domain.img_dir"), g.Cfg().GetString("Domain.img_dir"))
	s.AddStaticPath(g.Cfg().GetString("Domain.file_dir"), g.Cfg().GetString("Domain.file_dir"))
	//允许前端跨域
	GroupAPI := s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
	})

	//上传base64图片
	GroupAPI.ALL("/upload_img", controller.Upload_Base64)
	//上传文件
	GroupAPI.ALL("/upload_file", controller.Upload_File)
	s.Run()
}
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
