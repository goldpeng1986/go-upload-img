package tools

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"os"
)

type Config struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Bucket    string `yaml:"bucket"`
}

func Init(r *ghttp.Request) (*oss.Client, *Config, bool) {
	var mConfig *Config
	mConfig = &Config{
		Endpoint:  g.Cfg().GetString("Aliyun.Endpoint"),
		AccessKey: g.Cfg().GetString("Aliyun.AccessKeyId"),
		SecretKey: g.Cfg().GetString("Aliyun.AccessKeySecret"),
		Bucket:    "laozhu1986",
	}
	client, err := oss.New(mConfig.Endpoint, mConfig.AccessKey, mConfig.SecretKey)
	if err != nil {
		Return_Status(r, nil, "400", "初始化阿里云oss失败")
		return nil, nil, false
	}
	return client, mConfig, true
}
func Ali_Upload(r *ghttp.Request, mPath string) bool {
	mClient, mConfig, mBool := Init(r)
	if mBool == false {
		return false
	}
	mBucket, err := mClient.Bucket(mConfig.Bucket)
	if err != nil {
		Return_Status(r, nil, "400", "获取Bucket失败")
		return false
	}
	mDir, _ := os.Getwd()
	mDir = mDir + "/"
	// 检查文件是否存在
	if _, err := os.Stat(mDir + mPath); os.IsNotExist(err) {
		Return_Status(r, nil, "400", "文件不存在")
		return false
	}
	err = mBucket.PutObjectFromFile(mPath, mDir+mPath)
	if err != nil {
		Return_Status(r, nil, "400", "上传文件失败")
		return false
	}
	return true
}
