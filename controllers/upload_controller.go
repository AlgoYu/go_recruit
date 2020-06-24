package controllers

import (
	"cn.anydevelop/go_recruit/common"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"time"
)

const UPLOAD_DIRECTORY = "static/upload/"

// 上传图片
type UploadController struct {
	beego.Controller
}

// 上传图片
func (uploadController *UploadController) UploadPicture() {
	f, h, _ := uploadController.GetFile("picture") //获取上传的文件
	if f != nil || h != nil {
		ext := path.Ext(h.Filename)
		//验证后缀名是否符合要求
		var AllowExtMap = map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		}
		if _, ok := AllowExtMap[ext]; ok {
			//创建目录
			if _, err := os.Stat(UPLOAD_DIRECTORY); err != nil {
				os.MkdirAll(UPLOAD_DIRECTORY, os.ModePerm)
			}
			//构造文件名称
			rand.Seed(time.Now().UnixNano())
			randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
			hashName := md5.Sum([]byte(time.Now().Format("2020_06_20_15_04_05_") + randNum))

			fileName := fmt.Sprintf("%x", hashName) + ext

			fpath := UPLOAD_DIRECTORY + fileName
			defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
			err := uploadController.SaveToFile("picture", fpath)
			if err != nil {
				uploadController.Data["json"] = common.Fail(err.Error())
			} else {
				uploadController.Data["json"] = common.Success(fpath)
			}
		}
	} else {
		uploadController.Data["json"] = common.Fail("上传错误！")
	}
	uploadController.ServeJSON()
}
