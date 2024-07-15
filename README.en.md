# go-upload-img
#### Language
English | [中文](README.md)
#### Introduction
Deployment is very convenient and supports uploading images and various files to local storage, Alibaba Cloud OSS, Qiniu Cloud, and Tencent Cloud.
#### Software Architecture
Designed based on the Goframe framework, this module supports uploading images and files. It is recommended to deploy it separately and request it via POST or GET methods, facilitating secondary development in the future. Currently, local uploads and Alibaba Cloud OSS uploads are completed.
#### Installation tutorial
1. Download the code remotely, for example: `git clone https://gitee.com/gogokit/go-upload-img.git`
2. Set up the configuration file, for example: `config.toml`
3. Generate the binary file or exe file, for example: `go build -o ./linux_amd64/main main.go`
4. Deploy on BaoTa, configure the configuration file, and start it.
#### Usage Instructions
1. Access the address remotely or internally: `http://127.0.0.1:8080/api/upload_img`
2. Return result:
```
{
"code": "200",
"data": {
"img_dir": " http://127.0.0.1:8080/static/file/20240715/44.png ",
"img_name": "44.png"
},
"msg": "success"
}
```
3.  Done

#### Progress
| Module    | Status  | Notes |
|-------|-----|----|
| Image Upload  | Completed |    |
| File Upload  | Completed |    |
| Alibaba Cloud Upload | Completed |    |
| Qiniu Cloud Upload | Not Started |    |