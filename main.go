package main

import (
	"net/http"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
	"encoding/json"
	"time"
	"path/filepath"
)

//生成新文件名称
func generatorFileName(filenameType, originFileName string) (fileName string) {
	switch filenameType {
	//random暂时也用时间代替
	case "random":fallthrough
	case "datetime":
		datetime := time.Now().Format("20060102150405")
		fileName = fmt.Sprintf("%s%s",datetime,filepath.Ext(originFileName))
	default:
		fileName = originFileName
	}
	return
}

func put(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	endPoint := r.FormValue("endpoint")
	bucketName := r.FormValue("bucket")
	key := r.FormValue("key")
	secret := r.FormValue("secret")
	domain := r.FormValue("domain") //bucket域名
	filenameType := r.FormValue("filename_type") //文件名类型域名

	//新增client
	client, err := oss.New(endPoint, key, secret)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	//找到对应bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	//读取文件
	file, fileHandle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	defer file.Close()

	//新文件名称
	fileName := generatorFileName(filenameType, fileHandle.Filename)

	//上传文件
	err = bucket.PutObject(fileName, file)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	//返回文件url
	w.Header().Add("Content-Type", "application/json")
	data := map[string]string{"url":domain + "/" + fileName}
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", put)
	http.ListenAndServe(":8080", nil)
}