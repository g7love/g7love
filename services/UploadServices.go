package services

import (
	"fmt"
	"os"
	"io"
	"github.com/gin-gonic/gin"
	"time"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"math/rand"
)


//检查目录是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Print(filename + " not exist")
		exist = false
	}
	return exist
}
/**
 * User: lizhengxiang
 * Date: 17-5-28
 * Time: 上午8:55
 */
func Upload(c *gin.Context) string{
	name := c.PostForm("name")
	fmt.Println(name)
	file, header, err := c.Request.FormFile("UploadForm[imageFile]")
	if err != nil {
		return "Bad request"
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(time.Now().Unix(),10)+strconv.Itoa(r.Intn(100))))
	cipherStr := h.Sum(nil)
	filename := hex.EncodeToString(cipherStr)+header.Filename
	bool_fileexist := checkFileIsExist("./frontend/uploads")
	if bool_fileexist { //如果文件夹存在
		f, err := os.OpenFile("./frontend/uploads/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return "Bad request"
		}
		defer f.Close()
		io.Copy(f, file)
	} else {
		err1 := os.Mkdir("./frontend/uploads", os.ModePerm) //创建文件夹
		if err1 != nil {
			return "Bad request"
		}
		f, err := os.OpenFile("./frontend/uploads/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return "Bad request"
		}
		defer f.Close()
		io.Copy(f, file)
	}
	return "/uploads/"+filename
}