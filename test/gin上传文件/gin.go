package main

// 导入gin包
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/up", "./template")
	// 设置文件上传大小限制，默认是32m
	router.MaxMultipartMemory = 640 << 20 // 64 MiB

	router.POST("/upload", func(c *gin.Context) {
		// 获取上传文件，返回的是multipart.FileHeader对象，代表一个文件，里面包含了文件名之类的详细信息
		// file是表单字段名字
		file, _ := c.FormFile("file")
		// 打印上传的文件名
		log.Println(file.Filename)

		// 将上传的文件，保存到./data/1111.jpg 文件中
		c.SaveUploadedFile(file, "./data/1111.jpg")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8086")
}
