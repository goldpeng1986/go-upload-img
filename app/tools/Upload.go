package tools

import (
	"encoding/base64"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"io"
	"os"
	"strings"
)

// base64转图片 api版本
func Base64_To_Pic(r *ghttp.Request, mImg_base64 string, mImg_name string, mImg_path string) {
	//获取当前目录
	mDir, _ := os.Getwd()
	mDir = mDir + "/"
	//判断图片名称是否为空
	if g.IsEmpty(mImg_name) {
		mImg_name = Get_Ymdhis() + Get_Random_Int(4) + ".jpg" //默认jpg格式
	}
	//判断图片路径是否为空
	if g.IsEmpty(mImg_path) {
		mImg_path = g.Cfg().GetString("Domain.img_dir") + "/" + Get_Ymd() + "/"
	}
	// 提取Base64编码的数据
	mSplit := strings.Split(mImg_base64, ",")
	var mBase64Data string
	if len(mSplit) > 1 {
		mBase64Data = mSplit[1]
	} else {
		mBase64Data = mSplit[0]
	}
	// 解码Base64字符串为字节数组
	mByte_Data, err := base64.StdEncoding.DecodeString(mBase64Data)
	if err != nil {
		Return_Status(r, nil, "400", "解码Base64时出错,请确认图片格式是否正确")
		return
	}
	// 检查目录是否存在
	_, err = os.Stat(mDir + mImg_path)
	if os.IsNotExist(err) {
		// 目录不存在，执行创建目录的操作
		err := os.MkdirAll(mDir+mImg_path, 0755)
		if err != nil {
			Return_Status(r, nil, "400", "创建目录失败，请设置上级目录权限")
			return
		}
	}
	// 创建一个文件来保存图片
	mFiles := mDir + mImg_path + mImg_name
	file, err := os.Create(mFiles)
	if err != nil {
		Return_Status(r, nil, "400", "创建文件时出错，请设置上级目录权限")
		return
	}
	defer file.Close()

	// 将解码后的字节数组写入文件
	_, err = file.Write(mByte_Data)
	if err != nil {
		Return_Status(r, nil, "400", "写入二进制文件时出错，请检查图片是否正确")
		return
	}
	// 刷新文件缓存，确保数据被写入磁盘
	err = file.Sync()
	if err != nil {
		Return_Status(r, nil, "205", "刷新文件缓存时出错，请检查磁盘权限")
		return
	}
	//判断是否上传到阿里云
	if g.Cfg().GetBool("Aliyun.IsUpload") {
		mBool := Ali_Upload(r, mImg_path+mImg_name)
		if !mBool {
			return
		}
	}
	//返回图片地址和图片名称
	mReturn := g.Map{
		"img_dir":  g.Cfg().GetString("Domain.host") + mImg_path + mImg_name,
		"img_name": mImg_name,
	}
	Return_Status(r, mReturn, "", "")
	return
}

// 文件流转图片
func File_To_Pic(r *ghttp.Request, mFile *ghttp.UploadFile, mFile_name string, mFile_path string) {
	//fmt.Println(gjson.New(&mFile).Export())
	//打开文件流
	mFileReader, _ := mFile.Open()
	defer mFileReader.Close()
	//创建目录
	mDir, _ := os.Getwd() //获取绝对路径
	mDir = mDir + "/"
	//判断图片路径是否为空
	if g.IsEmpty(mFile_path) {
		mFile_path = g.Cfg().GetString("Domain.file_dir") + "/" + Get_Ymd() + "/"
	}
	//判断目录是否存在
	_, err := os.Stat(mDir + mFile_path)
	if err != nil {
		// 目录不存在，执行创建目录的操作
		err = gfile.Mkdir(mDir + mFile_path)
		if err != nil {
			Return_Status(r, nil, "400", "创建目录失败，请设置上级目录权限")
			return
		}
	}
	//判断文件名
	if g.IsEmpty(mFile_name) {
		mFile_name = mFile.Filename
	}
	// 打开文件以写入
	file, err := os.Create(mDir + mFile_path + mFile_name)
	if err != nil {
		Return_Status(r, nil, "400", "创建文件失败，请设置上级目录权限")
		return
	}
	defer file.Close()
	// 将文件流的内容写入文件
	_, err = io.Copy(file, mFileReader)
	if err != nil {
		Return_Status(r, nil, "400", "文件写入失败，请检查文件是否正确")
		return
	}
	//判断是否上传到阿里云
	if g.Cfg().GetBool("Aliyun.IsUpload") {
		mBool := Ali_Upload(r, mFile_path+mFile_name)
		if !mBool {
			return
		}
	}
	//返回图片地址和图片名称
	mReturn := g.Map{
		"file_dir":  g.Cfg().GetString("Domain.host") + mFile_path + mFile_name,
		"file_name": mFile_name,
	}
	Return_Status(r, mReturn, "", "")
	return
}
