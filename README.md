# go-upload-img

#### 介绍
部署非常方便，支持上传图片、各种文件模块，上传到本地和阿里云OSS、七牛云、腾讯云

#### 软件架构
基于Goframe框架设计的上传图片、文件模块，建议单独部署，通过post、get去请求，方便后期二次开发，目前已经完成本地上传和阿里云OSS上传


#### 安装教程

1.  远程下载代码，比如git clone https://gitee.com/gogokit/go-upload-img.git
2.  设置配置文件，比如config.toml
3.  生成二进制文件或者exe文件，比如go build -o ./linux_amd64/main main.go
4.  在宝塔上部署，配置好配置文件，然后启动即可

#### 使用说明

1.  远程或者内部访问地址：http://127.0.0.1:8080/api/upload_img
2.  返回结果：
```
 {
    "code": "200",
    "data": {
        "img_dir": "http://127.0.0.1:8080/static/file/20240715/44.png",
        "img_name": "44.png"
    },
    "msg": "success"
}
```
3.  完成

#### 进度
| 模块    | 状态  | 备注 |
|-------|-----|----|
| 上传图片  | 已完成 |    |
| 上传文件  | 已完成 |    |
| 上传阿里云 | 进行中 |    |
| 上传七牛云 | 未开始 |    |